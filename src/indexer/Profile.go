/*****************************************************************************
 *  file name : Profile.go
 *  author : Wu Yinghao
 *  email  : wyh817@gmail.com
 *
 *  file description : 正排索引基类，包括正排索引需要实现的接口
 *
******************************************************************************/

package indexer

import (
	u "utils"
)

const (
	PflNum = iota
	PflText
	PflDate
)

type Profile struct {
	Name string
	Type int64
	Len  int64
}



const FILT_TYPE_LESS  int64 = 1 
const FILT_TYPE_ABOVE int64 = 2 
const FILT_TYPE_EQUAL int64 = 3
const FILT_TYPE_UNEQUAL int64= 4 



type ProfileInterface interface {
	Put(doc_id int64, value interface{}) error
	Find(doc_id int64) (interface{}, error)
	Filter(doc_ids []u.DocIdInfo, value interface{}, is_forward bool,filt_type int64) ([]u.DocIdInfo, error)
	Display()
	GetType() int64
	GetMaxDocId() int64
	CustomFilter(doc_ids []u.DocIdInfo, value interface{}, r bool, cf func(v1, v2 interface{}) bool) ([]u.DocIdInfo, error)
}

/*****************************************************************************
*  function name : GetMaxDocId
*  params : nil
*  return : int64
*
*  description : get profile's length, max doc_id number
*
******************************************************************************/

func (this *Profile) GetMaxDocId() int64 {
	return this.Len - 1
}

/*****************************************************************************
*  function name : GetProfileName
*  params : nil
*  return : string
*
*  description : get profile name
*
******************************************************************************/

func (this *Profile) GetName() string {
	return this.Name
}
