package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/chromedp"
)

const (
	siteUrl = "https://www.avito.ru/avito-care/eco-impact"

	overridedRequest = "https://www.avito.ru/web/1/charity/ecoImpact/init"

	className = "desktop-impact-item-eeQO3"
)

func RunTest(request InitRequest, scFileName string) error {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Headless,
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Стартуем хром
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	chromedp.ListenTarget(ctx, func(event interface{}) {
		switch ev := event.(type) {
		case *fetch.EventRequestPaused:
			go func() {
				c := chromedp.FromContext(ctx)
				ctx := cdp.WithExecutor(ctx, c.Target)
				if ev.Request.URL == overridedRequest {
					fetch.FulfillRequest(ev.RequestID, int64(200)).
						WithBody(base64.StdEncoding.EncodeToString(body)).Do(ctx)
				}
				fetch.ContinueResponse(ev.RequestID).Do(ctx)
			}()
		}
	})

	// инициализируем пустой массив, куда будет сохранен скриншот
	var imageBuf []byte

	randTime := time.Duration(rand.Intn(10) + 7)

	// и отправляем хрому задачи, которые он должен выполнить
	// у нас только одна - ScreenshotTasks, но можно закинуть сколько угодно
	if err = chromedp.Run(
		ctx,
		fetch.Enable(),
		chromedp.Navigate(siteUrl),
		chromedp.Sleep(time.Second*randTime),
		chromedp.Screenshot(className, &imageBuf, chromedp.NodeVisible),
	); err != nil {
		return err
	}

	if err = os.WriteFile(scFileName, imageBuf, 0644); err != nil {
		return err
	}

	return nil
}
