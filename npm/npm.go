package npm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

import lib "github.com/DerYeger/npm-cards/lib"

func GetPackageData(packageName string, weeks int) (packageData lib.PackageData, err error) {
  packageData.Name = packageName
  for i := weeks; i >= 0; i-- {
    endDate := time.Now().AddDate(0, 0, -7 * i)
    startDate := time.Now().AddDate(0, 0, -7 * (i + 1))
    downloads, err := getDownloads(packageName, startDate, endDate)
    if err != nil {
      return packageData, err
    }
    packageData.WeeklyDownloads = append(packageData.WeeklyDownloads, downloads)
  }
  log.Print(packageData.WeeklyDownloads)
  return packageData, nil
}

func getDownloads(packageName string, startDate time.Time, endDate time.Time) (packageDownloads lib.PackageDownloads, err error) {

  endpoint := "https://api.npmjs.org/downloads/point/" + startDate.Format("2006-01-02") + ":" + endDate.Format("2006-01-02") + "/" + packageName
  log.Printf(endpoint)

  resp, err := http.Get(endpoint)
  if err != nil {
    return packageDownloads, err
  }
  if resp.StatusCode != 200 {
    return packageDownloads, fmt.Errorf("not found")
  }

  body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      return packageDownloads, err
   }

   err = json.Unmarshal(body, &packageDownloads)
   if err != nil {
    return packageDownloads, err
   }

   return packageDownloads, nil
}
