// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

import (
	"github.com/casbin/confita/util"
	"xorm.io/core"
)

type TreeItem struct {
	Key       string      `xorm:"varchar(100)" json:"key"`
	Title     string      `xorm:"varchar(100)" json:"title"`
	Content   string      `xorm:"mediumtext" json:"content"`
	TitleEn   string      `xorm:"varchar(100)" json:"titleEn"`
	ContentEn string      `xorm:"mediumtext" json:"contentEn"`
	Children  []*TreeItem `xorm:"varchar(1000)" json:"children"`
}

type Conference struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`

	Type           string   `xorm:"varchar(100)" json:"type"`
	Introduction   string   `xorm:"mediumtext" json:"introduction"`
	StartDate      string   `xorm:"varchar(100)" json:"startDate"`
	EndDate        string   `xorm:"varchar(100)" json:"endDate"`
	Organizer      string   `xorm:"varchar(100)" json:"organizer"`
	Carousels      []string `xorm:"mediumtext" json:"carousels"`
	CarouselHeight string   `xorm:"varchar(100)" json:"carouselHeight"`
	Tags           []string `xorm:"mediumtext" json:"tags"`

	DatasetUrl        string `xorm:"varchar(100)" json:"datasetUrl"`
	DatasetPreviewUrl string `xorm:"varchar(100)" json:"datasetPreviewUrl"`
	PreviewData       string `xorm:"mediumtext" json:"previewData"`
	ResultUrl         string `xorm:"varchar(100)" json:"resultUrl"`
	Bonus             int    `json:"bonus"`
	PersonCount       int    `json:"personCount"`
	DisplayState      string `xorm:"varchar(100)" json:"displayState"`

	Status           string      `xorm:"varchar(100)" json:"status"`
	Language         string      `xorm:"varchar(100)" json:"language"`
	Location         string      `xorm:"varchar(100)" json:"location"`
	Address          string      `xorm:"varchar(100)" json:"address"`
	EnableSubmission bool        `json:"enableSubmission"`
	DefaultItem      string      `xorm:"mediumtext" json:"defaultItem"`
	TreeItems        []*TreeItem `xorm:"mediumtext" json:"treeItems"`
}

func GetGlobalConferences() []*Conference {
	conferences := []*Conference{}
	err := adapter.engine.Asc("owner").Desc("created_time").Find(&conferences)
	if err != nil {
		panic(err)
	}

	return conferences
}

func GetConferences(owner string) []*Conference {
	conferences := []*Conference{}
	err := adapter.engine.Desc("created_time").Find(&conferences, &Conference{Owner: owner})
	if err != nil {
		panic(err)
	}

	return conferences
}

func getConference(owner string, name string) *Conference {
	conference := Conference{Owner: owner, Name: name}
	existed, err := adapter.engine.Get(&conference)
	if err != nil {
		panic(err)
	}

	if existed {
		return &conference
	} else {
		return nil
	}
}

func GetConference(id string) *Conference {
	owner, name := util.GetOwnerAndNameFromId(id)
	return getConference(owner, name)
}

func UpdateConference(id string, conference *Conference) bool {
	owner, name := util.GetOwnerAndNameFromId(id)
	if getConference(owner, name) == nil {
		return false
	}

	_, err := adapter.engine.ID(core.PK{owner, name}).AllCols().Update(conference)
	if err != nil {
		panic(err)
	}

	//return affected != 0
	return true
}

func AddConference(conference *Conference) bool {
	affected, err := adapter.engine.Insert(conference)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func DeleteConference(conference *Conference) bool {
	affected, err := adapter.engine.ID(core.PK{conference.Owner, conference.Name}).Delete(&Conference{})
	if err != nil {
		panic(err)
	}

	return affected != 0
}
