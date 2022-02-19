package router

import (
    "encoding/base64"
    "github.com/0xJacky/Nginx-UI/frontend"
    "github.com/0xJacky/Nginx-UI/server/model"
    "github.com/gin-contrib/static"
    "github.com/gin-gonic/gin"
    "io/fs"
    "log"
    "net/http"
    "path"
    "strings"
)

func authRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            tmp, _ := base64.StdEncoding.DecodeString(c.Query("token"))
            token = string(tmp)
            if token == "" {
                c.JSON(http.StatusForbidden, gin.H{
                    "message": "auth fail",
                })
                c.Abort()
                return
            }
        }

        n := model.CheckToken(token)

        if n < 1 {
            c.JSON(http.StatusForbidden, gin.H{
                "message": "auth fail",
            })
            c.Abort()
            return
        }
        c.Next()
    }
}

type serverFileSystemType struct {
    http.FileSystem
}

func (f serverFileSystemType) Exists(prefix string, _path string) bool {
    _, err := f.Open(path.Join(prefix, _path))
    return err == nil
}

func mustFS(dir string) (serverFileSystem static.ServeFileSystem) {

    sub, err := fs.Sub(frontend.DistFS, path.Join("dist", dir))

    if err != nil {
        log.Println(err)
        return
    }

    serverFileSystem = serverFileSystemType{
        http.FS(sub),
    }

    return
}

// tryStatic Static returns a middleware handler that serves static files in the given directory.
func tryStatic(urlPrefix string, fs static.ServeFileSystem) gin.HandlerFunc {
    fileserver := http.FileServer(fs)
    if urlPrefix != "" {
        fileserver = http.StripPrefix(urlPrefix, fileserver)
    }
    return func(c *gin.Context) {
        if fs.Exists(urlPrefix, c.Request.URL.Path) {
            fileserver.ServeHTTP(c.Writer, c.Request)
            if strings.Contains(c.Request.URL.Path, ".js") {
                c.Writer.Header().Set("content-type", "application/javascript")
            }
            c.Abort()
        }
    }
}