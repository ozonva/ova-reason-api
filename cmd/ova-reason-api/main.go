package main

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-reason-api/internal/flusher"
	"github.com/ozonva/ova-reason-api/internal/mocks"
	"github.com/ozonva/ova-reason-api/internal/model"

	. "github.com/onsi/ginkgo"
)

func main() {

	var (
		ctrl       *gomock.Controller
		mockRepo   *mocks.MockRepo
		reasons    []model.Reason
		flusherObj flusher.Flusher
	)
	RegisterFailHandler(Fail)
	ctrl = gomock.NewController(GinkgoT())
	mockRepo = mocks.NewMockRepo(ctrl)

	flusherObj = flusher.NewFlusher(2, mockRepo)
	reasons = []model.Reason{
		*model.New(1, 1, 1, "forgot keys"),
		*model.New(5, 5, 5, "lost my wallet"),
		*model.New(6, 6, 6, "want bonus"),
		*model.New(7, 7, 7, "my friend asked"),
		*model.New(8, 8, 8, "want to eat"),
	}

	mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil).Times(3)

	result := flusherObj.Flush(reasons)
	Expect(result).Should(BeNil())

	fmt.Println(result)

	gomock.InOrder(
		mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil),
		mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil),
		mockRepo.EXPECT().AddEntities(gomock.Any()).Return(errors.New("no enough space")),
	)

	result = flusherObj.Flush(reasons)

	fmt.Println(len(result))
	fmt.Println(result[0].String())
	Expect(result).ToNot(BeNil())
	Expect(len(result)).To(Equal(1))
	Expect(result[0]).To(Equal(reasons[4]))

}
