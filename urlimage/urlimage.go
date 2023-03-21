package urlimage

import (
	"context"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

type Options struct {
	Url      string
	Width    int
	Height   int
	Quality  int
	FilePath string
}

func Url2Image(option *Options) error {
	var (
		buf []byte
	)
	// 禁用 chrome headless
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoDefaultBrowserCheck,                        //不检查默认浏览器
		chromedp.Flag("headless", true),                       // 开启窗口模式
		chromedp.Flag("blink-settings", "imagesEnabled=true"), //开启图像界面,重点是开启这个
		chromedp.Flag("ignore-certificate-errors", true),      //忽略错误
		chromedp.Flag("disable-web-security", true),           //禁用网络安全标志
		chromedp.Flag("disable-extensions", true),             //开启插件支持
		chromedp.Flag("disable-default-apps", true),
		chromedp.WindowSize(option.Width, option.Height), // 设置浏览器分辨率（窗口大小）
		chromedp.Flag("disable-gpu", true),               //开启 gpu 渲染
		chromedp.Flag("hide-scrollbars", true),
		chromedp.Flag("mute-audio", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("no-default-browser-check", true),

		chromedp.NoFirstRun, //设置网站不是首次运行
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"), //设置UserAgent
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// capture entire browser viewport, returning png with quality=100
	if err := chromedp.Run(ctx, fullScreenshot(option.Url, 100, &buf)); err != nil {
		return err
	}
	if err := os.WriteFile(option.FilePath, buf, 0o644); err != nil {
		return err
	}
	return nil
}

func fullScreenshot(url string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(time.Second * 5),
		chromedp.FullScreenshot(res, quality),
	}
}
