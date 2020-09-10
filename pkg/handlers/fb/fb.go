package fb

import (
	cgstorage "cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
	"fmt"
	"google.golang.org/api/option"
	"io"
	"mime/multipart"
)

const (
	SERVICE_ACCOUNT_KEY = "fir-test-6bee5-firebase-adminsdk-8wt0f-c98261114e.json"
	PROJECT_NAME        = "fir-test-6bee5"
)

type Fb struct {
	app           *firebase.App
	storageClient *storage.Client
	auth          *auth.Client
	context       context.Context
}

func (fb *Fb) getClient() error {
	if fb.storageClient == nil {
		config := &firebase.Config{
			StorageBucket: PROJECT_NAME + ".appspot.com",
		}
		opt := option.WithCredentialsFile(SERVICE_ACCOUNT_KEY)
		var err error
		fb.context = context.Background()
		fb.app, err = firebase.NewApp(fb.context, config, opt)
		if err != nil {
			return err
		}
		fb.storageClient, err = fb.app.Storage(fb.context)
		if err != nil {
			return err
		}
		fb.auth, err = fb.app.Auth(fb.context)
		if err != nil {
			return err
		}
	}
	return nil
}

func (fb *Fb) Upload(filename string, srcFile multipart.File) (string, error) {
	if err := fb.getClient(); err != nil {
		return "", err
	}
	bucket, err := fb.storageClient.DefaultBucket()
	if err != nil {
		return "", err
	}
	obj := bucket.Object(filename)
	wc := obj.NewWriter(fb.context)
	if _, err := io.Copy(wc, srcFile); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	if err := obj.ACL().Set(fb.context, cgstorage.AllUsers, cgstorage.RoleReader); err != nil {
		return "", err
	}
	//attrs, err := obj.Attrs(fb.context)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(attrs.MediaLink)
	//fmt.Println("https://storage.googleapis.com/" + PROJECT_NAME + ".appspot.com/" + filename)
	return "https://storage.googleapis.com/" + PROJECT_NAME + ".appspot.com/" + filename, nil
}

func (fb *Fb) GetUsers() (*[]string, error) {
	if err := fb.getClient(); err != nil {
		return nil, err
	}
	//fb.auth.GetUsers(fb.context)
	userRecord, err := fb.auth.GetUserByEmail(fb.context, "timurkash@gmail.com")
	if err != nil {
		return nil, err
	}
	fmt.Println(userRecord.UID)
	customClaims := map[string]interface{}{"isAdmin": true}

	fmt.Println(userRecord)
	if err := fb.auth.SetCustomUserClaims(fb.context, userRecord.UID, customClaims); err != nil {
		return nil, err
	}
	users := []string{}
	return &users, nil
}
