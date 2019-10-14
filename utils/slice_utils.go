package utils

import (
	"fmt"

	"github.com/SerhiiCho/shoshka-go/models"
)

// Contains function returns true if slice contains given item
func Contains(slice []string, needle string) bool {
	for _, sliceItem := range slice {
		if sliceItem == needle {
			return true
		}
	}
	return false
}

// GetUniqueItem returns unique item from a slice
func GetUniqueItem(slice1 []string, slice2 []string) []string {
	var result []string

	for _, item1 := range slice1 {
		if !Contains(slice2, item1) {
			result = append(result, item1)
		}
	}

	for _, item2 := range slice2 {
		if !Contains(slice1, item2) {
			result = append(result, item2)
		}
	}

	return result
}

// GetIndexOfSliceItem returns index of a slice item where value is positioned
func GetIndexOfSliceItem(slice []string, value string) int {
	for i, sliceItem := range slice {
		if sliceItem == value {
			return i
		}
	}

	return -1
}

// GenerateMapOfNewData returns nil if there are no new photo reports
func GenerateMapOfNewData(titles []string, photoReports []models.PhotoReport) []models.PhotoReport {
	oldTitles := GetCached("titles")
	newPhotoReportTitles := GetUniqueItem(titles, oldTitles)

	if len(titles) > 0 {
		PutIntoCache(titles, "titles")
	}

	if len(newPhotoReportTitles) == 0 {
		fmt.Print("There are no new photo reports")
		return nil
	}

	var tgMessageData []models.PhotoReport

	for _, newReportTitle := range newPhotoReportTitles {
		index := GetIndexOfSliceItem(titles, newReportTitle)

		if index == -1 {
			continue
		}

		tgMessageData = append(tgMessageData, models.PhotoReport{
			Title: newReportTitle,
			Image: photoReports[index].Image,
			URL:   photoReports[index].URL,
		})
	}

	return tgMessageData
}

// GetTitlesFromPhotoReports returns only titles from slice of messsage data
func GetTitlesFromPhotoReports(photoReports []models.PhotoReport) []string {
	var titles []string

	for _, photoReport := range photoReports {
		titles = append(titles, photoReport.Title)
	}

	return titles
}
