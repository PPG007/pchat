package utils

import (
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"pchat/repository/bson"
	"testing"
	"time"
)

type TestA struct {
	Id            bson.ObjectId
	CreatedAt     time.Time
	B             *TestB
	CaseSensitive string
}

type TestB struct {
	Id            string
	CreatedAt     string
	Casesensitive string
	Id2           string
}

func TestObjectIdStringConverter(t *testing.T) {
	a := TestA{
		Id: bson.NewObjectId(),
	}
	b := TestB{}
	err := DefaultCopier().From(a).To(&b)
	assert.NoError(t, err)
	assert.Equal(t, a.Id.Hex(), b.Id)
}

func TestStringObjectIdConverter(t *testing.T) {
	a := TestA{}
	b := TestB{Id: bson.NewObjectId().Hex()}
	err := DefaultCopier().From(b).To(&a)
	assert.NoError(t, err)
	assert.Equal(t, b.Id, a.Id.Hex())
}

func TestTimeStringConverter(t *testing.T) {
	a := TestA{CreatedAt: time.Now()}
	b := TestB{}
	err := DefaultCopier().From(a).To(&b)
	assert.NoError(t, err)
	assert.Equal(t, a.CreatedAt.Format(time.RFC3339), b.CreatedAt)
}

func TestStringTimeConverter(t *testing.T) {
	a := TestA{}
	b := TestB{CreatedAt: time.Now().Format(time.RFC3339)}
	err := DefaultCopier().From(b).To(&a)
	assert.NoError(t, err)
	assert.Equal(t, b.CreatedAt, a.CreatedAt.Format(time.RFC3339))
}

func TestCopier_CaseSensitive(t *testing.T) {
	a := TestA{CaseSensitive: bson.NewObjectId().Hex()}
	b := TestB{}
	err := DefaultCopier().From(a).CaseSensitive().To(&b)
	assert.NoError(t, err)
	assert.NotEqual(t, a.CaseSensitive, b.Casesensitive)
}

func TestCopier_DeepCopy(t *testing.T) {
	a := TestA{
		B: &TestB{
			Id: bson.NewObjectId().Hex(),
		},
	}
	aa := TestA{}
	err := DefaultCopier().From(a).DeepCopy().To(&aa)
	assert.NoError(t, err)
	a.B.Id = bson.NewObjectId().Hex()
	assert.NotEqual(t, a.B.Id, aa.B.Id)
}

func TestCopier_RegisterDiffPair(t *testing.T) {
	a := TestA{Id: bson.NewObjectId()}
	b := TestB{}
	err := DefaultCopier().RegisterDiffPair([]copier.FieldNameMapping{
		{
			SrcType: TestA{},
			DstType: TestB{},
			Mapping: map[string]string{
				"Id": "Id2",
			},
		},
	}).From(a).To(&b)
	assert.NoError(t, err)
	assert.Equal(t, a.Id.Hex(), b.Id2)
}
