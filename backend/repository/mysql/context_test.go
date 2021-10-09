package mysql

import (
	"testing"
	"time"

	"github.com/Thewalkers2012/DOJ/model"
	"github.com/Thewalkers2012/DOJ/util/random"
	"github.com/stretchr/testify/assert"
)

func createRandomConetext(t *testing.T) *model.Context {
	req := &model.CreateContextParams{
		Title:     random.RandomUserName(),
		Author:    random.RandomUserName(),
		StartTime: model.Time(time.Now()),
		EndTime:   model.Time(time.Now()),
	}

	context, err := CreateContext(req)
	assert.NoError(t, err)
	assert.Equal(t, req.Title, context.Title)
	assert.Equal(t, req.Author, context.Author)
	assert.Equal(t, req.StartTime, context.StartTime)
	assert.Equal(t, context.EndTime, req.EndTime)

	assert.NotZero(t, context.ID)

	return context
}

func TestCraeteContext(t *testing.T) {
	createRandomConetext(t)
}

func TestGetContextByID(t *testing.T) {
	context1 := createRandomConetext(t)

	req := &model.GetContextParams{
		ContextID: context1.ID,
	}

	context2, err := GetContextByID(req.ContextID)

	assert.NoError(t, err)
	assert.Equal(t, context1.ID, context2.ID)
	assert.Equal(t, context1.Author, context2.Author)
	assert.Equal(t, context1.Defunct, context2.Defunct)
	assert.WithinDuration(t, time.Time(context1.StartTime), time.Time(context2.StartTime), time.Second)
	assert.WithinDuration(t, time.Time(context1.EndTime), time.Time(context2.EndTime), time.Second)
}
