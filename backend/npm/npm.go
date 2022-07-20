package npm

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/DerYeger/npm-cards/backend/lib"
	"golang.org/x/sync/errgroup"
)

func GetPackageData(packageName string, weeks int) (packageData lib.PackageData, err error) {
  packageData.Name = packageName
  eg := errgroup.Group{}
  mu := &sync.Mutex{}

  for i := weeks; i >= 0; i-- {
    startDate := time.Now().AddDate(0, 0, -7 * (i + 1)).Format("2006-01-02")
    endDate := time.Now().AddDate(0, 0, -7 * i).Format("2006-01-02")
    eg.Go(func() error {
      endpoint := "https://api.npmjs.org/downloads/point/" + startDate + ":" + endDate + "/" + packageName
      resp, err := http.Get(endpoint)
      if err != nil {
        return err
      }
      if resp.StatusCode != 200 {
        return errors.New("404")
      }

      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
          return err
      }
      var packageDownloads lib.PackageDownloads
      err = json.Unmarshal(body, &packageDownloads)
      if err != nil {
        return err
      }
      mu.Lock()
      packageData.WeeklyDownloads = append(packageData.WeeklyDownloads, packageDownloads)
      mu.Unlock()
      return nil
    })
  }

  err = eg.Wait()
  if err != nil {
    return packageData, err
  }

  sort.Slice(packageData.WeeklyDownloads, func(i, j int) bool {
    return packageData.WeeklyDownloads[i].Start < packageData.WeeklyDownloads[j].Start
  })

  return packageData, nil
}
