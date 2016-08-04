package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/meinside/opengov-go/types"

	_ "github.com/mattn/go-sqlite3" // $ go get -u github.com/mattn/go-sqlite3
)

type SaveType string

const (
	TypeSqlite SaveType = "sqlite"

	// TODO - support more types
)

const (
	SqliteFilename = "opengov.sqlite"
)

// 저장
func save(saveType SaveType, infoList []types.InfoList, publicList []types.PublicList, researchList []types.ResearchList) error {
	var err error

	switch saveType {
	case TypeSqlite:
		err = saveSqlite(infoList, publicList, researchList)
	}

	return err
}

// .sqlite 파일로 저장
func saveSqlite(infoList []types.InfoList, publicList []types.PublicList, researchList []types.ResearchList) error {
	log.Printf("> Saving to sqlite...")

	var dbFilepath string
	if currentDir, err := os.Getwd(); err != nil {
		return err
	} else {
		dbFilepath = filepath.Join(currentDir, SqliteFilename)
	}

	if db, err := sql.Open("sqlite3", dbFilepath); err == nil {
		// drop tables if exist
		if _, err := db.Exec(`drop table if exists info_lists`); err != nil {
			return err
		}
		if _, err := db.Exec(`drop table if exists public_lists`); err != nil {
			return err
		}
		if _, err := db.Exec(`drop table if exists research_lists`); err != nil {
			return err
		}

		// create tables
		if _, err := db.Exec(`create table info_lists(
			id integer primary key autoincrement,
			package_id text not null,
			doc_prdctn_dt text not null,
			trck_card_nm text not null,
			title text not null,
			src_dept_doc_id text not null,
			writer text not null,
			othnd_pd text not null,
			dept_nm text not null,
			othbs_se text not null,
			cpyrht text,
			url text
		)`); err != nil {
			return err
		}
		if _, err := db.Exec(`create table public_lists(
			id integer primary key autoincrement,
			nid text not null,
			category text not null,
			title text not null,
			writer text not null,
			dept_nm text not null,
			regist_dt text not null,
			taxonomy text not null,
			telno text,
			cpyrht text not null,
			url text not null
		)`); err != nil {
			return err
		}
		if _, err := db.Exec(`create table research_lists(
			id integer primary key autoincrement,
			nid text not null,
			title text not null,
			regist_dt text not null,
			relm_cl text not null,
			creat_yr text not null,
			category text not null,
			region text not null,
			isbn text,
			relte_area text not null,
			writer text not null,
			doc_prdctn_dt text not null,
			cpyrht text not null,
			othbs_se text not null,
			job_se text not null,
			url text not null
		)`); err != nil {
			return err
		}

		// TODO - create indices

		// insert values to each table
		for _, info := range infoList {
			if stmt, err := db.Prepare(`insert into info_lists(
				package_id,
				doc_prdctn_dt,
				trck_card_nm,
				title,
				src_dept_doc_id,
				writer,
				othnd_pd,
				dept_nm,
				othbs_se,
				cpyrht,
				url
			) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`); err != nil {
				return err
			} else {
				defer stmt.Close()
				if _, err = stmt.Exec(
					info.PackageId,
					info.DocumentProductionDate,
					info.TrackCardName,
					info.Title,
					info.SourceDepartmentDocumentId,
					info.Writer,
					info.OthndPeriod,
					info.DepartmentName,
					info.OthbsSe,
					info.Copyright,
					info.Url,
				); err != nil {
					return err
				}
			}
		}
		for _, public := range publicList {
			if stmt, err := db.Prepare(`insert into public_lists(
				nid,
				category,
				title,
				writer,
				dept_nm,
				regist_dt,
				taxonomy,
				telno,
				cpyrht,
				url
			) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`); err != nil {
				return err
			} else {
				defer stmt.Close()
				if _, err = stmt.Exec(
					public.Nid,
					public.Category,
					public.Title,
					public.Writer,
					public.DepartmentName,
					public.RegisterDate,
					public.Taxonomy,
					public.TelephonNumber,
					public.Copyright,
					public.Url,
				); err != nil {
					return err
				}
			}
		}
		for _, research := range researchList {
			if stmt, err := db.Prepare(`insert into research_lists(
				nid,
				title,
				regist_dt,
				relm_cl,
				creat_yr,
				category,
				region,
				isbn,
				relte_area,
				writer,
				doc_prdctn_dt,
				cpyrht,
				othbs_se,
				job_se,
				url
			) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`); err != nil {
				return err
			} else {
				defer stmt.Close()
				if _, err = stmt.Exec(
					research.Nid,
					research.Title,
					research.RegisterDate,
					research.RelmCl,
					research.CreateYear,
					research.Category,
					research.Region,
					research.Isbn,
					research.RelteArea,
					research.Writer,
					research.DocumentProductionDate,
					research.Copyright,
					research.OthbsSe,
					research.JobSe,
					research.Url,
				); err != nil {
					return err
				}
			}
		}
	} else {
		return err
	}

	log.Printf("> Saved to: %s\n", dbFilepath)

	return nil
}
