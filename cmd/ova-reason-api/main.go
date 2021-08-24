package main

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-reason-api/internal/flusher"
	"github.com/ozonva/ova-reason-api/internal/mocks"
	"github.com/ozonva/ova-reason-api/internal/model"
	"github.com/ozonva/ova-reason-api/internal/saver"
	"strconv"
	"sync"
	"time"

	. "github.com/onsi/ginkgo"
)

func main() {

	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo

		flusherObj flusher.Flusher
		saverObj   saver.Saver
	)
	RegisterFailHandler(Fail)
	ctrl = gomock.NewController(GinkgoT())
	mockRepo = mocks.NewMockRepo(ctrl)

	flusherObj = flusher.NewFlusher(2, mockRepo)

	mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil).Times(100000)

	saverObj = saver.NewSaver(5, flusherObj)
	wg := sync.WaitGroup{}
	wg.Add(3)

	go runGenerator("goroutine1", 1000, saverObj, &wg)
	go runGenerator("goroutine2", 1500, saverObj, &wg)
	go runGenerator("goroutine3", 1800, saverObj, &wg)

	wg.Wait()
	saverObj.Close()

}

func runGenerator(id string, timeout int, saverObj saver.Saver, wg *sync.WaitGroup) {
	cnt := 0
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * time.Duration(timeout))
		cnt++
		newReason := *model.New(1, 1, 1, id+" "+strconv.Itoa(cnt))
		saverObj.Save(newReason)

	}
	wg.Done()
}
