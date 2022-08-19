package util

import (
	"regexp"
	"strings"
	"github.com/is0405/hacku/model"
)

func CheckUser(mu *model.User, passwordCheck bool) bool {
	if mu.Name == "" || !NameCheck(mu.Name) {
		return false
	}
	mu.Name = ReplaceString(mu.Name)
	
	if mu.Age < 18 {
		return false
	}

	//法(2)・産社(2)・国際(2)・文(2)・言語・先端・映像(2)・経済(2)・スポーツ(2)・食マネ(2)・理工(2)・情理(2)・生命(2)・薬(2)・経営(2)・政策(2)・心理・グローバル・人間科学・テク
	if mu.Faculty < 0 && 20 <= mu.Faculty {
		return false
	}

	
	if mu.Gender < 0 && 3 <= mu.Gender {
		return false
	}

	if mu.Mail == "" || !MailCheck(mu.Mail) {
		return false
	}
	mu.Mail = ReplaceString(mu.Mail)

	if passwordCheck {
		if mu.Password == "" {
			return false
		}

		if strings.Contains(mu.Password, "`") {
			return false
		}
	}
	
	return true
}

func CheckRecruitment(ma *model.Recruitment) bool {
	if ma.Contents == "" {
		return false
	}
	ma.Contents = ReplaceString(ma.Contents)

	if ma.Conditions == "" {
		return false
	}
	ma.Conditions = ReplaceString(ma.Conditions)

	if ma.Reward == "" {
		return false
	}
	ma.Reward = ReplaceString(ma.Reward)

	if ma.MaxParticipation <= 0 {
		return false
	}

	if ma.StartRecruitmentPeriod == "" {
		return false
	}
	ma.StartRecruitmentPeriod = ReplaceString(ma.StartRecruitmentPeriod)

	if ma.FinishRecruitmentPeriod == "" {
		return false
	}
	ma.FinishRecruitmentPeriod = ReplaceString(ma.FinishRecruitmentPeriod)

	if ma.StartImplementationPeriod == "" {
		return false
	}
	ma.StartImplementationPeriod = ReplaceString(ma.StartImplementationPeriod)

	if ma.FinishImplementationPeriod == "" {
		return false
	}
	ma.FinishImplementationPeriod = ReplaceString(ma.FinishImplementationPeriod)

	if ma.Title == "" {
		return false
	}
	ma.Title = ReplaceString(ma.Title)

	if ma.Gender < 0 || 3 < ma.Gender {
		return false
	}

	if ma.MinAge == -1 {
		ma.MinAge = 18
	}

	if ma.MaxAge == -1 {
		ma.MaxAge = 60
	}

	if ma.MaxAge < ma.MinAge || ma.MinAge < 0 || 100 < ma.MaxAge {
		return false
	}

	return true
}

func MailCheck(str string) bool {
	chars := []string{"@", "."}
    r := strings.Join(chars, "")
	symbol := regexp.MustCompile("[^" + r + "A-Za-z0-9]+")
	if symbol.Match([]byte(str)) {
		//上記以外がある
		return false
	} else {
		symbol := regexp.MustCompile(`\s*@\s*`)
		symbol2 := regexp.MustCompile(`\s*\.\s*`)

		group := symbol.Split(str, -1)
		if len(group) != 2 {
			return false
		}

		group = symbol2.Split(str, -1)
		for i := 0; i < len(group); i++ {
			if group[i] == "" {
				return false
			} else if strings.HasSuffix(group[i], "@") {
				return false
			}
		}
	}
	return true
}

func NameCheck(str string) bool {
	chars := []string{"?", "!", "\\*","\\_", "\\#", "<", ">", "\\\\", "(", ")", "\\$", "\"", "%", "=", "~", "|", "[", "]", ";", "\\+", ":", "{", "}", "@", "\\`", "/", "；", "＠", "＋", "：", "＊", "｀", "「", "」", "｛", "｝", "＿", "？", "。", "、", "＞", "＜"}
    r := strings.Join(chars, "")
	symbol := regexp.MustCompile("[" + r + "]+")
	if symbol.Match([]byte(str)) {
		//上記が含まれている
		return false
	}
	return true
}

func ReplaceString(str string) string {
	str = strings.Replace(str, "`", "'", -1)
	return str
}
