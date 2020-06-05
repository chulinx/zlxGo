package color

//输出带颜色的字符串
type Color struct {
}

// newColor
func NewColor() *Color {
	return &Color{}
}

//红色
func (color Color) Red(text string) string {
	return "\033[31m" + text + "\033[0m" //输出红色
}

//绿色
func (color Color) Green(text string) string {
	return "\033[32m" + text + "\033[0m"
}

// 黄色

func (color Color)Yellow(text string) string  {
	return "\033[33m" + text + "\033[0m"
}
//下划线
func (color Color) Underline(text string) string {
	return "\033[4m" + text + "\033[0m"
}