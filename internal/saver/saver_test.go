package saver_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-reason-api/internal/flusher"
	"github.com/ozonva/ova-reason-api/internal/saver"
	"time"

	"github.com/ozonva/ova-reason-api/internal/mocks"
	"github.com/ozonva/ova-reason-api/internal/model"
)

var _ = Describe("Saver", func() {

	var (
		ctrl       *gomock.Controller
		mockRepo   *mocks.MockRepo
		reasons    []model.Reason
		flusherObj flusher.Flusher
		saverObj   saver.Saver
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)

		reasons = []model.Reason{
			*model.New(1, 1, 1, "forgot keys"),
			*model.New(5, 5, 5, "lost my wallet"),
			*model.New(6, 6, 6, "want bonus"),
			*model.New(7, 7, 7, "my friend asked"),
			*model.New(8, 8, 8, "want to eat"),
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Using normal capacity", func() {
		BeforeEach(func() {
			flusherObj = flusher.NewFlusher(2, mockRepo)
			saverObj = saver.NewSaver(10, flusherObj)
		})

		It("Saving after Close", func() {
			savedCounter := 0

			mockRepo.EXPECT().AddEntities(gomock.Any()).DoAndReturn(func(arg []model.Reason) interface{} {
				savedCounter += len(arg)
				return nil
			}).AnyTimes()

			for _, reason := range reasons {
				saverObj.Save(reason)
			}

			//before Close
			Expect(savedCounter).Should(Equal(0))

			saverObj.Close()

			//after Close
			Expect(savedCounter).Should(Equal(len(reasons)))
		})

		It("Saving by timeout", func() {
			counter := 0

			mockRepo.EXPECT().AddEntities(gomock.Any()).DoAndReturn(func(arg []model.Reason) interface{} {
				counter += len(arg)
				return nil
			}).AnyTimes()

			for _, reason := range reasons {
				saverObj.Save(reason)
			}

			//before Sleep
			Expect(counter).Should(Equal(0))

			time.Sleep(5 * time.Second)

			//after Sleep
			Expect(counter).Should(Equal(len(reasons)))
		})

		It("Double Close", func() {
			savedCounter := 0

			mockRepo.EXPECT().AddEntities(gomock.Any()).DoAndReturn(func(arg []model.Reason) interface{} {
				savedCounter += len(arg)
				return nil
			}).AnyTimes()

			for _, reason := range reasons {
				saverObj.Save(reason)
			}

			//before Close
			Expect(savedCounter).Should(Equal(0))

			saverObj.Close()

			//after Close
			Expect(savedCounter).Should(Equal(len(reasons)))

			saverObj.Close()

			//after second Close
			Expect(savedCounter).Should(Equal(len(reasons)))
		})
	})

	Describe("Using low capacity", func() {
		BeforeEach(func() {
			flusherObj = flusher.NewFlusher(2, mockRepo)
			saverObj = saver.NewSaver(1, flusherObj)
		})

		It("Saving after Close", func() {
			savedCounter := 0

			mockRepo.EXPECT().AddEntities(gomock.Any()).DoAndReturn(func(arg []model.Reason) interface{} {
				savedCounter += len(arg)
				return nil
			}).AnyTimes()

			for _, reason := range reasons {
				saverObj.Save(reason)
			}

			//before Close
			Expect(savedCounter).Should(Equal(0))

			saverObj.Close()

			//after Close
			Expect(savedCounter).Should(Equal(len(reasons)))
		})

		It("Saving by timeout", func() {
			counter := 0

			mockRepo.EXPECT().AddEntities(gomock.Any()).DoAndReturn(func(arg []model.Reason) interface{} {
				counter += len(arg)
				return nil
			}).AnyTimes()

			for _, reason := range reasons {
				saverObj.Save(reason)
			}

			//before Sleep
			Expect(counter).Should(Equal(0))

			time.Sleep(5 * time.Second)

			//after Sleep
			Expect(counter).Should(Equal(len(reasons)))
		})
	})

})
