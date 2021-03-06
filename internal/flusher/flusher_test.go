package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-reason-api/internal/flusher"
	"github.com/ozonva/ova-reason-api/internal/mocks"
	"github.com/ozonva/ova-reason-api/internal/model"
)

var _ = Describe("Flusher", func() {

	var (
		ctrl       *gomock.Controller
		mockRepo   *mocks.MockRepo
		reasons    []model.Reason
		flusherObj flusher.Flusher
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

	Describe("Flushing slice of reasons into base using chunks", func() {
		BeforeEach(func() {
			flusherObj = flusher.NewFlusher(2, mockRepo)
		})
		Context("With no error", func() {

			It("should return null", func() {
				gomock.InOrder(
					mockRepo.EXPECT().AddEntities(gomock.Len(2)).Return(nil),
					mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil),
					mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil),
				)
				result := flusherObj.Flush(reasons)
				Expect(result).Should(BeNil())
			})
		})

		Context("With errors", func() {

			It("should return the last entity", func() {
				gomock.InOrder(
					mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil),
					mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil),
					mockRepo.EXPECT().AddEntities(gomock.Any()).Return(errors.New("no enough space")),
				)
				result := flusherObj.Flush(reasons)
				Expect(result).ToNot(BeNil())
				Expect(len(result)).To(Equal(1))
				Expect(result[0]).To(Equal(reasons[4]))

			})
		})
	})

	Describe("Flushing slice of reasons into base one by one", func() {
		BeforeEach(func() {
			flusherObj = flusher.NewFlusher(1, mockRepo)
		})
		Context("With no error", func() {

			It("should return null", func() {

				mockRepo.EXPECT().AddEntities(gomock.Len(1)).Return(nil).Times(5)

				result := flusherObj.Flush(reasons)
				Expect(result).Should(BeNil())
			})
		})

		Context("With errors", func() {

			It("should return the third entity", func() {
				gomock.InOrder(
					mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil),
					mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil),
					mockRepo.EXPECT().AddEntities(gomock.Any()).Return(errors.New("duplicated keys")),
					mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil),
					mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil),
				)
				result := flusherObj.Flush(reasons)
				Expect(result).ToNot(BeNil())
				Expect(len(result)).To(Equal(1))
				Expect(result[0]).To(Equal(reasons[2]))

			})
		})
	})
})
