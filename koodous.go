package main

import ( 
  "encoding/json"
  "github.com/parnurzeal/gorequest"
  "log"
  "fmt"
  "os"
  "net/http"
  "io"
  "bytes"
  "mime/multipart"
  "time"
  "flag"
)

var APIKEY = ""

// APKInfo JSON Struct 
type apkinfo struct {
    CreatedOn int `json:"created_on"`
    Rating int `json:"rating"`
    Image string `json:"image"`
    Tags []interface{} `json:"tags"`
    Md5 string `json:"md5"`
    Sha1 string `json:"sha1"`
    Sha256 string `json:"sha256"`
    App string `json:"app"`
    PackageName string `json:"package_name"`
    Company string `json:"company"`
    DisplayedVersion string `json:"displayed_version"`
    Size int `json:"size"`
    Stored bool `json:"stored"`
    Analyzed bool `json:"analyzed"`
    IsApk bool `json:"is_apk"`
    Trusted bool `json:"trusted"`
    Detected bool `json:"detected"`
    Corrupted bool `json:"corrupted"`
    Repo string `json:"repo"`
    OnDevices bool `json:"on_devices"`
}


// Userinfo JSON Struct
type userinfo struct {
  AvatarURL string `json:"avatar_url"`
  DateJoined int `json:"date_joined"`
  LastLogin int `json:"last_login"`
  TotalPublicRulesets int `json:"total_public_rulesets"`
  TotalComments int `json:"total_comments"`
  IsSuperuser bool `json:"is_superuser"`
  Username string `json:"username"`
  FirstName string `json:"first_name"`
  LastName string `json:"last_name"`
  Occupation string `json:"occupation"`
  Bio string `json:"bio"`
  TwitterUser string `json:"twitter_user"`
  TotalFollowers int `json:"total_followers"`
  TotalFollowing int `json:"total_following"`
  TotalSocialDetections int `json:"total_social_detections"`
  Latest24HSocialDetections int `json:"latest_24h_social_detections"`
  TotalVotes int `json:"total_votes"`
  Following []interface{} `json:"following"`
}

// Votes struct 
type votes_inf struct {
  Count int `json:"count"`
  Next interface{} `json:"next"`
  Previous interface{} `json:"previous"`
  Results []struct {
    Kind string `json:"kind"`
    Analyst string `json:"analyst"`
  } `json:"results"`
}

// Comments struct 

type comments_info struct {
  Next interface{} `json:"next"`
  Previous interface{} `json:"previous"`
  Results []struct {
    ID int `json:"id"`
    Author struct {
      AvatarURL string `json:"avatar_url"`
      DateJoined int `json:"date_joined"`
      LastLogin int `json:"last_login"`
      TotalPublicRulesets int `json:"total_public_rulesets"`
      TotalComments int `json:"total_comments"`
      IsSuperuser bool `json:"is_superuser"`
      Username string `json:"username"`
      FirstName string `json:"first_name"`
      LastName string `json:"last_name"`
      Occupation interface{} `json:"occupation"`
      Bio interface{} `json:"bio"`
      TwitterUser interface{} `json:"twitter_user"`
      TotalFollowers int `json:"total_followers"`
      TotalFollowing int `json:"total_following"`
      TotalSocialDetections int `json:"total_social_detections"`
      Latest24HSocialDetections int `json:"latest_24h_social_detections"`
      TotalVotes int `json:"total_votes"`
      Following []interface{} `json:"following"`
    } `json:"author"`
    CreatedOn int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
    Apk string `json:"apk"`
    Ruleset interface{} `json:"ruleset"`
    Text string `json:"text"`
  } `json:"results"`
}

// Public Ruleset struct
type public_ruleset struct {
  ID int `json:"id"`
  CreatedOn int `json:"created_on"`
  ModifiedOn int `json:"modified_on"`
  Analyst struct {
    AvatarURL string `json:"avatar_url"`
    DateJoined int `json:"date_joined"`
    LastLogin int `json:"last_login"`
    TotalPublicRulesets int `json:"total_public_rulesets"`
    TotalComments int `json:"total_comments"`
    IsSuperuser bool `json:"is_superuser"`
    Username string `json:"username"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Occupation interface{} `json:"occupation"`
    Bio string `json:"bio"`
    TwitterUser string `json:"twitter_user"`
    TotalFollowers int `json:"total_followers"`
    TotalFollowing int `json:"total_following"`
    TotalSocialDetections int `json:"total_social_detections"`
    Latest24HSocialDetections int `json:"latest_24h_social_detections"`
    TotalVotes int `json:"total_votes"`
    Following []int `json:"following"`
  } `json:"analyst"`
  Name string `json:"name"`
  Rules string `json:"rules"`
  Active bool `json:"active"`
  Privacy string `json:"privacy"`
  Social bool `json:"social"`
  PendingSocial bool `json:"pending_social"`
  Deleted bool `json:"deleted"`
  SendNotifications bool `json:"send_notifications"`
  Detections int `json:"detections"`
  Rating int `json:"rating"`
  Parent interface{} `json:"parent"`
}

type ruleset_list struct {
  Count int `json:"count"`
  Next string `json:"next"`
  Previous interface{} `json:"previous"`
  Results []struct {
    ID int `json:"id"`
    CreatedOn int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
    Analyst struct {
      AvatarURL string `json:"avatar_url"`
      DateJoined int `json:"date_joined"`
      LastLogin int `json:"last_login"`
      TotalPublicRulesets int `json:"total_public_rulesets"`
      TotalComments int `json:"total_comments"`
      IsSuperuser bool `json:"is_superuser"`
      Username string `json:"username"`
      FirstName string `json:"first_name"`
      LastName string `json:"last_name"`
      Occupation interface{} `json:"occupation"`
      Bio interface{} `json:"bio"`
      TwitterUser interface{} `json:"twitter_user"`
      TotalFollowers int `json:"total_followers"`
      TotalFollowing int `json:"total_following"`
      TotalSocialDetections int `json:"total_social_detections"`
      Latest24HSocialDetections int `json:"latest_24h_social_detections"`
      TotalVotes int `json:"total_votes"`
      Following []interface{} `json:"following"`
    } `json:"analyst"`
    Name string `json:"name"`
    Rules string `json:"rules"`
    Active bool `json:"active"`
    Privacy string `json:"privacy"`
    Social bool `json:"social"`
    PendingSocial bool `json:"pending_social"`
    Deleted bool `json:"deleted"`
    SendNotifications bool `json:"send_notifications"`
    Detections int `json:"detections"`
    Rating int `json:"rating"`
    Parent interface{} `json:"parent"`
  } `json:"results"`
}

// Obtain userinfo given an user 
func user_info(user string) {
  request := gorequest.New()
  _, body, _ := request.Get("https://api.koodous.com/analysts/" + user).
  Set("Authorization", APIKEY).
  End()
   
  var user_details userinfo
  byteArray := []byte(body)
  json.Unmarshal(byteArray, &user_details)

  fmt.Println("Displaying user information ... ")
  fmt.Println("========================================")
  fmt.Println(user_details.Username + " - \""+ user_details.Bio +"\"")
  fmt.Println("Social Detections:", user_details.TotalSocialDetections)
  fmt.Println("Total rulesets:", user_details.TotalPublicRulesets)
  fmt.Println("Comments:", user_details.TotalComments, "\t", "Votes:", user_details.TotalVotes)
  fmt.Println("Following:", user_details.TotalFollowing, "\t", "Followers:", user_details.TotalFollowers)
  t := time.Unix(int64(user_details.DateJoined), 0 )
  fmt.Println("Joined:", t)
  fmt.Println("========================================")
}

/* Given a SHA256 the whole JSON object */
func get_apk_info(checksum string) {
  request := gorequest.New()
  _, body, _ := request.Get("https://api.koodous.com/apks/" + checksum).
  Set("Authorization", APIKEY).
  End()

  var apk_info apkinfo 
  byteArray := []byte(body)
  json.Unmarshal(byteArray, &apk_info)

  fmt.Println("Displaying ", apk_info.PackageName + " ...")
  fmt.Println("====================================================================")
  fmt.Println("Author: " + apk_info.Company + "\t Size:" , apk_info.Size)
  if apk_info.Analyzed == true {
    fmt.Println("This application has already been scanned")
  } else {
    fmt.Println("Aw! Looks like it's not been scanned. Wait for Koodous to analyze it or request analysis.")
  }
  if apk_info.Detected == true {
    fmt.Println("This app has been detected as a threat by Koodous' Community.")
  } else {
    fmt.Println("This app is clean. If you think it's malware, report it!")
  }
  if apk_info.OnDevices == true {
    fmt.Println("WARNING! This app is installed on devices!")
  }
  fmt.Println("Rating:" , apk_info.Rating,  "\t" , "Trusted?:" , apk_info.Trusted)
  fmt.Println("====================================================================")
  fmt.Println("SHA256:" , apk_info.Sha256)
  fmt.Println("SHA1:" , apk_info.Sha1)
  fmt.Println("MD5:" , apk_info.Md5)
  fmt.Println("====================================================================")
  fmt.Println("You can check this apk at https://www.koodous.com/apks/" + checksum)
}

func downloadFromUrl(url string, checksum string) {
  fileName := checksum
  fmt.Println("Downloading", url, "to", fileName)

  // TODO: check file existence first with io.IsExist
  output, err := os.Create(fileName)
  if err != nil {
    fmt.Println("Error while creating", fileName, "-", err)
    return
  }
  defer output.Close()

  response, err := http.Get(url)
  if err != nil {
    fmt.Println("Error while downloading", url, "-", err)
    return
  }
  defer response.Body.Close()

  n, err := io.Copy(output, response.Body)
  if err != nil {
    fmt.Println("Error while downloading", url, "-", err)
    return
  }

  fmt.Println(n, "bytes downloaded.")
}

func Upload(url, file string) (err error) {
    // Prepare a form that you will submit to that URL.
    var b bytes.Buffer
    w := multipart.NewWriter(&b)
    // Add your image file
    f, err := os.Open(file)
    if err != nil {
        return
    }
    defer f.Close()
    fw, err := w.CreateFormFile("image", file)
    if err != nil {
        return
    }
    if _, err = io.Copy(fw, f); err != nil {
        return
    }
    // Add the other fields
    if fw, err = w.CreateFormField("key"); err != nil {
        return
    }
    if _, err = fw.Write([]byte("KEY")); err != nil {
        return
    }
    // Don't forget to close the multipart writer.
    // If you don't close it, your request will be missing the terminating boundary.
    w.Close()

    // Now that you have a form, you can submit it to your handler.
    req, err := http.NewRequest("POST", url, &b)
    if err != nil {
        return
    }
    // Don't forget to set the content type, this will contain the boundary.
    req.Header.Set("Content-Type", w.FormDataContentType())

    // Submit the request
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return
    }

    // Check the response
    if res.StatusCode != http.StatusOK {
        err = fmt.Errorf("bad status: %s", res.Status)
    }
    return
}


/* Given a local file (SHA256) it will be uploaded
to Koodous in order to be analyzed */ 
func upload_apk(checksum string) {
  request := gorequest.New() 
  _, res, _ := request.Get("https://api.koodous.com/apks/" + checksum + "/get_upload_url").
  Set("Authorization", "Token " + APIKEY).
  End()
  var f interface{}
  byteArray := []byte(res)
  err := json.Unmarshal(byteArray, &f)

  if err != nil {
    log.Fatal(err)
  }

  m := f.(map[string]interface{})

  for _, v := range m {
    Upload(v.(string), checksum)
  }
}


/* Looks up at Koodous the checksum given, then generates
a temporary download URL and retrieves the file. */ 
func get_apk(checksum string) {
  request := gorequest.New()
  _, res, _ := request.Get("https://api.koodous.com/apks/" + checksum + "/download").
  Set("Authorization", "Token " + APIKEY).
  End()
  var f interface{}
  byteArray := []byte(res)
  err := json.Unmarshal(byteArray, &f)
  if err != nil {
    log.Fatal(err)
  }
  m := f.(map[string]interface{})

  for _, v := range m {
    downloadFromUrl(v.(string), checksum)
    fmt.Println("File uploaded!")
  }
}

// Obtain votes of given SHA 
func retrieve_votes(checksum string) {
  request := gorequest.New() 
  _, res, _ := request.Get("https://api.koodous.com/apks/" + checksum + "/votes").
  Set("Authorization", "Token " + APIKEY).
  End() 

  var apk_votes votes_inf
  byteArray := []byte(res)
  json.Unmarshal(byteArray, &apk_votes)
  pos := 0
  neg := 0 
  fmt.Println("\t\tTotal votes:", len(apk_votes.Results))
  for i:=0; i < len(apk_votes.Results); i++ {
    if apk_votes.Results[i].Kind == "negative" {
      neg++
    } else {
      pos++
    }
  }
  fmt.Println("POSITIVE: ", pos, "\t", "NEGATIVE: ", neg)
}

// Self explanatory
func vote_up(checksum string) {
  request := gorequest.New() 
  request.Post("https://api.koodous.com/apks/" + checksum + "/votes").
  Set("Authorization", "Token " + APIKEY).
  Send(`{"kind": "positive"}`).
  End()
  fmt.Println("Positive vote issued")
}

// Self explanatory
func vote_down(checksum string) {
  request := gorequest.New() 
  request.Post("https://api.koodous.com/apks/" + checksum + "/votes").
  Set("Authorization", "Token " + APIKEY).
  Send(`{"kind": "negative"}`).
  End()
  fmt.Println("Negative vote issued")
}

// Retrieve complete list of comments
func retrieve_comments(checksum string) {
  request := gorequest.New() 
  _, res, _ := request.Get("https://api.koodous.com/apks/" + checksum + "/comments").
  Set("Authorization", "Token " + APIKEY).
  End()

  byteArray := []byte(res)
  var comments comments_info  
  json.Unmarshal(byteArray, &comments)
  for i := 0; i < len(comments.Results); i++ {
      t := time.Unix(int64(comments.Results[i].CreatedOn), 0 )
    fmt.Println(t)
    fmt.Println(comments.Results[i].ID, "> " + comments.Results[i].Author.Username + ": \"" + comments.Results[i].Text + "\"")
  }
}

// Create comment 
func comment_create(checksum string, comment string) {
  request := gorequest.New()
  _, res , _ := request.Post("https://api.koodous.com/apks/" + checksum + "/comments").
  Set("Authorization", "Token " + APIKEY).
  Send(`{"text": "`+ comment + `"}`).
  End()
  fmt.Println(res)

}

// Delete comment ID. 
func comment_delete(checksum string, id string) {
  request := gorequest.New()
  _, res , _ := request.Delete("https://api.koodous.com/apks/" + checksum + "/comments/" + id).
  Set("Authorization", "Token " + APIKEY).
  End()
  fmt.Println(res)

}

// Get given ruleset
func get_ruleset(id string) {
  request := gorequest.New() 
  _, res, _ := request.Get("https://api.koodous.com/public_rulesets/" + id).
  Set("Authorization", "Token " + APIKEY).
  End()
  byteArray := []byte(res)
  var ruleset public_ruleset
  json.Unmarshal(byteArray, &ruleset)
  if ruleset.Name == "" {
    fmt.Println("Either rule doesn't exist or you are not able to access it. Private maybe?")
  } else {
    fmt.Println("Ruleset: " + ruleset.Name + " (" + ruleset.Analyst.Username + ")")
    if ruleset.Social == true {
      fmt.Println("Social? Yes!" +  " Rating:", ruleset.Rating)
    } else {
      fmt.Println("Social? No" +  " Rating:", ruleset.Rating)
    }
    fmt.Println()
  }
}

/* Retrieve a page of rulesets, including information such as 
name, number of detections, author ... */

func get_ruleset_list(page string) {
  request := gorequest.New() 
  _, res, _ := request.Get("https://api.koodous.com/public_rulesets?page=" + page + "&page_size=10&active=True&privacy=public&ordering=-modified_on").
  Set("Authorization", "Token " + APIKEY).
  End()
  byteArray := []byte(res)
  var ruleset_page ruleset_list
  json.Unmarshal(byteArray, &ruleset_page)
  for i,_ := range ruleset_page.Results {
    fmt.Println("=============================")
    fmt.Println("\t\t" + ruleset_page.Results[i].Analyst.Username)
    fmt.Println("ID:", ruleset_page.Results[i].ID, "\t Name:" + ruleset_page.Results[i].Name)
    if ruleset_page.Results[i].Active == true {
      fmt.Println("Detections:", ruleset_page.Results[i].Detections, "\tActive? Yes!")    
    } else {
      fmt.Println("Detections:", ruleset_page.Results[i].Detections, "\tActive? No")
    }
    if ruleset_page.Results[i].Social == true {
      fmt.Println("Social? Yes!" +  " Rating:", ruleset_page.Results[i].Rating)
    } else {
      fmt.Println("Social? No" +  " Rating:", ruleset_page.Results[i].Rating)
    }
    fmt.Println("=============================")
  }
}

func main() {

  // Help 
  help := `
  * getuser <username>: Returns user info.
  * apk <sha256>: Returns APK info.
  * ruleset <id>: Returns ruleset content.
  * listrulesets <page-number>: Returns complete ruleset page. 
  * votes <sha256>: Returns given votes given a hash.
  * comments <sha256>: Returns comments posted on a given hash.
  * voteup <sha256>: Vote up given hash.
  * votedown <sha256>: Downvote given hash.
  * createcomment <sha256> <comment>: Post comment
  * deletecomment <sha256> <commentid>: Delete comment.
  * getapk <sha256>: Download APK
  * uploadapk <filepath>: Upload APK
  `

  /* Checking whether the argument is valid, 
  in case not, a list of commands with help will be displayed. */ 

  flag.Parse()
  args := flag.Args()

  if len(args) == 0 {
    // Nothing to do here.
    fmt.Println(help)

  } else {
    switch args[0] {
      case "getuser":
        user_info(args[1])
      case "apk":
        get_apk_info(args[1])
      case "ruleset":
        get_ruleset(args[1])
      case "comments":
        retrieve_comments(args[1])
      case "votes":
        retrieve_votes(args[1])
      case "votedown":
        vote_down(args[1])
      case "voteup":
        vote_up(args[1])
      case "listrulesets":
        get_ruleset_list(args[1])
      case "createcomment":
        comment_create(args[1], args[2])
      case "deletecomment":
        comment_delete(args[1], args[2])
      case "version":
        fmt.Println(".::   .::                           .::                         ")
        fmt.Println(".::  .::                            .::                         ")
        fmt.Println(".:: .::       .::       .::         .::   .::    .::  .:: .:::: ")
        fmt.Println(".: .:       .::  .::  .::  .::  .:: .:: .::  .:: .::  .::.::    ")
        fmt.Println(".::  .::   .::    .::.::    .::.:   .::.::    .::.::  .::  .::: ")
        fmt.Println(".::   .::   .::  .::  .::  .:: .:   .:: .::  .:: .::  .::    .::")
        fmt.Println(".::     .::   .::       .::     .:: .::   .::      .::.::.:: .::")
        fmt.Println("Version 0.1.0")
      default:
        fmt.Print(help)
    }
  }
}