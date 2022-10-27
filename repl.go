package main

import (
	"os"
	"regexp"
	"strings"
	"sync"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"git.lolli.tech/lollipopkit/lk/api"
	"git.lolli.tech/lollipopkit/lk/consts"
	"git.lolli.tech/lollipopkit/lk/state"
)

var (
	linesHistory  = []string{}
	blockStartReg = regexp.MustCompile(strings.Join([]string{
		consts.ForInReStr,
		consts.FnReStr,
		consts.WhileReStr,
		consts.IfReStr,
		consts.ElseIfReStr,
		consts.ElseReStr,
		consts.ClassDefReStr,
	}, "|"))
	blockEndReg = regexp.MustCompile("} *$")
	promptLen   = len([]rune(prompt))
	printReg    = regexp.MustCompile(`print\(.*\)`)
)

const (
	prompt = "➜ "
)

func repl(wg *sync.WaitGroup) {
	ls := state.New()
	ls.OpenLibs()
	println("REPL - Lang LK (v" + consts.VERSION + ")")

	blockStr := ""
	blockStartCount := 0
	blockEndCount := 0
	wg.Wait()
	for {
		os.Stdout.WriteString(prompt)

		line := readline()
		if line == "" {
			continue
		}
		if blockStartReg.MatchString(line) {
			blockStartCount++
		}
		if blockEndReg.MatchString(line) {
			blockEndCount++
		}

		blockStr += line
		if blockStartCount != blockEndCount {
			blockStr += "\n"
		}

		cmd := ""
		if blockStartCount > 0 && blockStartCount == blockEndCount {
			cmd = blockStr
		} else if blockStartCount > 0 {
			continue
		} else {
			blockStr = ""
			cmd = line
		}
		// println("==", cmd, "==")

		// 更新历史记录
		updateHistory(cmd)

		// 加载line，调用
		protectedLoadString(&ls, cmd)
		ls.PCall(0, -1, 0, true)
		blockStartCount = 0
		blockEndCount = 0
		blockStr = ""
	}
}

func catchErr(ls *api.LkState, first *bool, cmd string) {
	if err := recover(); err != nil {
		defer catchErr(ls, first, cmd)
		if *first {
			*first = false
			(*ls).LoadString(cmd, "stdin")
		}
	}
}

func protectedLoadString(ls *api.LkState, cmd string) {
	first := true
	// 捕获错误
	defer catchErr(ls, &first, cmd)
	addedPrintCmd := func() string {
		if printReg.MatchString(cmd) {
			return cmd
		}
		return "print(" + cmd + ")"
	}()
	(*ls).LoadString(addedPrintCmd, "stdin")
}

func _updateHistory(str string) {
	idx := -1
	for i := range linesHistory {
		if linesHistory[i] == str {
			idx = i
			break
		}
	}
	if idx != -1 {
		linesHistory = append(linesHistory[:idx], linesHistory[idx+1:]...)
	}
	linesHistory = append(linesHistory, str)
}

func updateHistory(str string) {
	str = strings.Trim(str, "\n")
	strs := strings.Split(str, "\n")
	for idx := range strs {
		_updateHistory(strs[idx])
	}
}

func readline() string {
	str := ""
	linesIdx := len(linesHistory)
	cursorIdx := 0

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.CtrlC, keys.Escape:
			os.Exit(0)
		case keys.RuneKey:
			runes := key.Runes
			s := string(runes)
			str = str[:cursorIdx] + s + str[cursorIdx:]
			cursorIdx += len(s)
			resetLine(str)
		case keys.Enter:
			println()
			return true, nil
		case keys.Backspace:
			if len(str) > 0 && cursorIdx > 0 {
				str = str[:cursorIdx-1] + str[cursorIdx:]
				resetLine(str)
				cursorIdx--
			}
		case keys.Left:
			if cursorIdx > 0 {
				cursorIdx--
			}
		case keys.Right:
			if cursorIdx < len(str) {
				cursorIdx++
			}
		case keys.Up:
			if linesIdx > 0 {
				linesIdx--
				str = linesHistory[linesIdx]
				resetLine(str)
				cursorIdx = len(str)
			}
		case keys.Down:
			if linesIdx < len(linesHistory)-1 {
				linesIdx++
				str = linesHistory[linesIdx]
				resetLine(str)
				cursorIdx = len(str)
			} else if linesIdx == len(linesHistory)-1 {
				str = ""
				resetLine("")
				cursorIdx = 0
			}
		case keys.Space:
			str += " "
			print(" ")
			cursorIdx++
		case keys.Tab:
			str += "  "
			print("  ")
			cursorIdx += 2
		case keys.Delete:
			if cursorIdx < len(str) {
				str = str[:cursorIdx] + str[cursorIdx+1:]
				resetLine(str)
			}
		}

		cursor.HorizontalAbsolute(cursorIdx + promptLen)
		return false, nil
	})
	return str
}

func resetLine(str string) {
	cursor.ClearLine()
	cursor.StartOfLine()
	print(prompt + str)
}
