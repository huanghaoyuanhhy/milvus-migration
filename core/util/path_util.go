package util

import (
	"errors"
	"github.com/zilliztech/milvus-migration/core/type/milvustype"
	"path/filepath"
	"strconv"
	"strings"
)

func GetSourceRVFilePath(tablesDir string, segColInfo *milvustype.SegColInfo) string {
	return filepath.Join(tablesDir, segColInfo.CollectionName, segColInfo.SegmentName, segColInfo.SegmentName+".rv")
}

func GetSourceUIDFilePath(tablesDir string, segColInfo *milvustype.SegColInfo) string {
	return filepath.Join(tablesDir, segColInfo.CollectionName, segColInfo.SegmentName, segColInfo.SegmentName+".uid")
}

func GetSourceDeletedDocsFilePath(tablesDir string, segColInfo *milvustype.SegColInfo) string {
	return filepath.Join(tablesDir, segColInfo.CollectionName, segColInfo.SegmentName, "deleted_docs")
}

func GetOutputUIDFilePath(outputDir string, segColInfo *milvustype.SegColInfo) (string, string) {
	targetOutputDir := outputDir
	targetDir := filepath.Join(targetOutputDir, "tables", segColInfo.CollectionName, segColInfo.SegmentName)
	targetFileName := filepath.Join(targetDir, "id.npy")

	return targetDir, targetFileName
}

func GetOutputRVFilePath(outputDir string, segColInfo *milvustype.SegColInfo) (string, string) {
	targetOutputDir := outputDir
	targetDir := filepath.Join(targetOutputDir, "tables", segColInfo.CollectionName, segColInfo.SegmentName)
	targetFileName := filepath.Join(targetDir, "data.npy")

	return targetDir, targetFileName
}

func GetOutputMetaJsonFilePath(outputDir string) (string, string) {
	outputMetaJson := filepath.Join(outputDir, "meta.json")
	return outputDir, outputMetaJson
}

func GenerateFaissIdFilePath(outputDir string, colName string) (string, string) {
	targetDir := filepath.Join(outputDir, colName)
	fileName := filepath.Join(targetDir, "id.npy")
	return targetDir, fileName
}

func GenerateFaissDataFilePath(outputDir string, colName string) (string, string) {
	targetDir := filepath.Join(outputDir, colName)
	fileName := filepath.Join(targetDir, "data.npy")
	return targetDir, fileName
}

func GenerateESDataFilePath(outputDir string, indexName string) string {
	targetDir := filepath.Join(outputDir, indexName)
	return targetDir
	//fileName := filepath.Join(targetDir, "data.json")
	//return targetDir, fileName
}

func GenerateESDataSubFileName(targetDir string, subTaskNum int, sort int) string {
	//data_1_1.json, data_1_2.json , data_2_3.json, data_2_4.json
	fileNmae := strings.Join([]string{"data", strconv.Itoa(subTaskNum), strconv.Itoa(sort)}, "_")
	fileNmae = strings.Join([]string{fileNmae, "json"}, ".")
	return filepath.Join(targetDir, fileNmae)
}

func GetAddressAndPortFromEndpoint(endpoint string) (string, string, error) {
	if endpoint == "" {
		return "", "", errors.New("endpoint cannot empty")
	}

	split := strings.Split(endpoint, ":")
	return split[0], split[1], nil
}
