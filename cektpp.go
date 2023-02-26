package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "log"
    "time"
    "regexp"
    // "github.com/bitly/go-simplejson"
    // "github.com/thedevsaddam/gojsonq"
)

type ErrorResponse struct {
    Error string `json:"error"`
}


func handleResponse(resp *http.Response) error {
    var errResp ErrorResponse
    err := json.NewDecoder(resp.Body).Decode(&errResp)
    if err != nil {
        return fmt.Errorf("failed to decode response body: %v", err)
    }

    if resp.StatusCode >= 400 {
        return fmt.Errorf("received error response: %d %s", resp.StatusCode, errResp.Error)
    }

    // handle successful response
    return nil
}


// type itemdata struct {
//     data1 string // I have tried making these strings
//     data2 string
//     data3 string
// }

type itemdata [][]string
type Bird struct {
    Species string
    Description string
}

type Root struct {
    Kdunit string
    Nmunit string
}

type Librari struct {
    Root []struct {
        Nama_program string `json:"NMPRGRM"`
        Nama_unit    string `json:"NMUNIT"`
        Ket    string `json:"KET"`
    } `json:"root"`
}

func check(result string  ) bool {
    // faulty regex   
    // m, err := regexp.MatchString("b\\ello w\\b",result)
    m, err := regexp.MatchString("Penyediaan Gaji dan Tunjangan ASN",result)
    if err != nil {
      fmt.Println("your regex is faulty")
      // you should log it or throw an error 
    //   return err.Error()
    }
    if (m) {
        // fmt.Println("Found it ")
        return true
    } else {
        return false
    }
}

const jsonText = `
{
  "libraries": [{
    "name": "Katakrak",
    "city": "Pamplona",
    "country": "Spain",
    "books": [{
            "title": "Harry Potter and the Sorcerer's Stone",
            "author": "J. K. Rowling",
            "availability": true
            }]
    },
    {
    "name": "Re-Read Librería Lowcost",
    "city": "Pamplona",
    "country": "Spain",
    "books": [{
            "title": "The Subtle Art of Not Giving a F*ck",
            "author": "Mark Manson",
            "availability": false
            }]
    }
]}`

func main() {
    // resp, err := http.Get("https://echoof.me/json")
    resp, err := http.Get("http://e-sp2d.bandaacehkota.go.id/storeweb.php?_dc=1677312100805&SKPD=DINAS%20KOMUNIKASI%2C%20INFORMATIKA%20DAN%20STATISTIK&NOSPM=&REKTUJUAN=&NAMAKEGIATAN=&limit=30&token=tcOVdreOvLaPtr%2FDkna3t52Hi6%2B3yZmCpcK6soOBtbuVVbfCsruLtMXDkliH&start=NaN")

    if err != nil {
      // handle error
      // fmt.Printf("failed to %s\n", err)
      // writeAPI.WritePoint(context.Background(), p)
      time.Sleep(10 * time.Second)
    }else{
        defer resp.Body.Close()
        // err = handleResponse(resp)

        // if err != nil {
        //     fmt.Println(err)
        // }else{
            b, _ := io.ReadAll(resp.Body)
          
            // var res map[string]interface{}
            var res map[string]any
            jsonErr := json.Unmarshal(b, &res)
            if jsonErr != nil {
                log.Fatal(jsonErr)
            }else{
                // var datas itemdata
                // settings, _ := simplejson.NewJson(b)
                // // json.Unmarshal(res["root"], &datas)
                // // fmt.Printf("\nJson: %+v",settings)
                // // fmt.Printf("\nJson: %+v",settings)
                // for k, v := range settings.MustMap() {
                //     fmt.Println("\n k:",k)
                //     fmt.Println("\n v:",v)
                // }

                var librariesInformation Librari
                err := json.Unmarshal(b, &librariesInformation)
                if err != nil {
                    log.Fatal("error unmarshaling json: ", err)
                }
                
                // log.Printf("librariesInformation: %+v", librariesInformation.Root)


                // birdJson := `{"root":[{"NO":1,"KDUNIT":"2.16.2.20.2.21.01.00.","NMUNIT":"DINAS KOMUNIKASI, INFORMATIKA DAN STATISTIK","NMPRGRM":"Penyediaan Jasa Komunikasi, Sumber Daya Air dan Listrik","NOSPM":"012\/2.16.2.20.2.21.01.00\/SPM\/LS\/2023","NOSP2D":"00393\/SP2D\/LS\/2023","REKTUJUAN":"PT. ACEHLINK MEDIA","WKTEVENT":"23-02-2023, 15:21:54","KET":"SP2D Telah Cair"},{"NO":2,"KDUNIT":"2.16.2.20.2.21.01.00.","NMUNIT":"DINAS KOMUNIKASI, INFORMATIKA DAN STATISTIK","NMPRGRM":"Penyediaan Gaji dan Tunjangan ASN","NOSPM":"009\/2.16.2.20.2.21.01.00\/SPM\/LS\/2023","NOSP2D":"00247\/SP2D\/LS\/2023","REKTUJUAN":"Bendahara SKPD","WKTEVENT":"13-02-2023, 10:41:35","KET":"SP2D Telah Cair"},{"NO":3,"KDUNIT":"2.16.2.20.2.21.01.00.","NMUNIT":"DINAS KOMUNIKASI, INFORMATIKA DAN STATISTIK","NMPRGRM":"Penyediaan Gaji dan Tunjangan ASN","NOSPM":"007\/2.16.2.20.2.21.01.00\/SPM\/LS\/2023","NOSP2D":"00246\/SP2D\/LS\/2023","REKTUJUAN":"Bendahara SKPD","WKTEVENT":"13-02-2023, 10:36:29","KET":"SP2D Telah Cair"},{"NO":4,"KDUNIT":"2.16.2.20.2.21.01.00.","NMUNIT":"DINAS KOMUNIKASI, INFORMATIKA DAN STATISTIK","NMPRGRM":"Penyediaan Jasa Komunikasi, Sumber Daya Air dan Listrik","NOSPM":"008\/2.16.2.20.2.21.01.00\/SPM\/LS\/2023","NOSP2D":"00210\/SP2D\/LS\/2023","REKTUJUAN":"PERANTARA TRS KASDA OL - KU SP2D","WKTEVENT":"10-02-2023, 11:31:29","KET":"SP2D Telah Cair"},{"NO":5,"KDUNIT":"2.16.2.20.2.21.01.00.","NMUNIT":"DINAS KOMUNIKASI, INFORMATIKA DAN STATISTIK","NMPRGRM":"Penyediaan Gaji dan Tunjangan ASN","NOSPM":"006\/2.16.2.20.2.21.01.00\/SPM\/LS\/2023","NOSP2D":"00170\/SP2D\/LS\/2023","REKTUJUAN":"Bendahara SKPD","WKTEVENT":"03-02-2023, 10:39:53","KET":"SP2D Telah Cair"},{"NO":6,"KDUNIT":"2.16.2.20.2.21.01.00.","NMUNIT":"DINAS KOMUNIKASI, INFORMATIKA DAN STATISTIK","NMPRGRM":"Penyediaan Gaji dan Tunjangan ASN","NOSPM":"005\/2.16.2.20.2.21.01.00\/SPM\/LS\/2023","NOSP2D":"00109\/SP2D\/LS\/2023","REKTUJUAN":"Bendahara SKPD","WKTEVENT":"30-01-2023, 11:42:49","KET":"SP2D Telah Cair"},{"NO":7,"KDUNIT":"2.16.2.20.2.21.01.00.","NMUNIT":"DINAS KOMUNIKASI, INFORMATIKA DAN STATISTIK","NMPRGRM":"Penyediaan Administrasi Pelaksanaan Tugas ASN","NOSPM":"004\/2.16.2.20.2.21.01.00\/SPM\/LS\/2023","NOSP2D":"00061\/SP2D\/LS\/2023","REKTUJUAN":"Bendahara SKPD","WKTEVENT":"26-01-2023, 14:28:42","KET":"SP2D Telah Cair"},{"NO":8,"KDUNIT":"2.16.2.20.2.21.01.00.","NMUNIT":"DINAS KOMUNIKASI, INFORMATIKA DAN STATISTIK","NMPRGRM":"UANG PERSEDIAAN","NOSPM":"002\/2.16.2.20.2.21.01.00\/SPM\/UP\/2023","NOSP2D":"00006\/SP2D\/UP\/2023","REKTUJUAN":"Bendahara SKPD","WKTEVENT":"11-01-2023, 16:01:00","KET":"SP2D Telah Cair"}],"total":"8"}`
                // var result map[string]any
                // json.Unmarshal([]byte(birdJson), &result)

                // root := result["root"].([]map[string]any)
                // fmt.Println(result)

                for _, data := range librariesInformation.Root {
                    // fmt.Printf("%s (%d)\n", person.Name, person.Age)
                    if check(data.Nama_program) {
                        fmt.Printf("%s, %s %s \n", data.Nama_unit, data.Nama_program, data.Ket)
                    }
                    
                    // fmt.Println(data.Nama_unit)
                }


                // for key, value := range librariesInformation.Root {
                //     // Each value is an `any` type, that is type asserted as a string
                //     fmt.Println(key, value.(string))
                // }

                // birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
                // var birds []Bird
                // json.Unmarshal([]byte(birdJson), &birds)
                // fmt.Printf("Birds : %+v", birds)

                // var settings map[string]interface{}
                // if err := json.NewDecoder(resp.Body).Decode(&settings); err != nil {
                //     panic(err)
                // }
                // fmt.Printf("\nJson: %+v",settings)

                // fmt.Println(settings)
                // fmt.Println(b)
                // fmt.Printf("%s\n", res["root"])
        //         // fmt.Printf("Voltage 2 : %0.1f\n", res["voltage2"].(float64))

        //         // voltage1 = res["voltage1"].(float64)
            }
        // } // end handle err response
    } // end if err http
    
}
