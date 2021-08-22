package main

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-reason-api/internal/flusher"
	"github.com/ozonva/ova-reason-api/internal/mocks"
	"github.com/ozonva/ova-reason-api/internal/model"
	"github.com/ozonva/ova-reason-api/internal/saver"
	"strconv"
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

	mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil).Times(100000) //.Do(func() { fmt.})

	saverObj = saver.NewSaver(5, flusherObj)

	go runGenerator("goroutine1", 1000, saverObj)
	go runGenerator("goroutine2", 1500, saverObj)
	go runGenerator("goroutine3", 1800, saverObj)
	time.Sleep(20 * time.Second)

	saverObj.Close()

}

func runGenerator(id string, timeout int, saverObj saver.Saver) {
	cnt := 0
	for {
		time.Sleep(time.Millisecond * time.Duration(timeout))
		cnt++
		newReason := *model.New(1, 1, 1, id+" "+strconv.Itoa(cnt))
		saverObj.Save(newReason)
	}
}
