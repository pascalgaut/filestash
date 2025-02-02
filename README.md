![screenshot](https://raw.githubusercontent.com/mickael-kerjean/filestash_images/master/.assets/photo.jpg)

<p align="center">
    <a href="https://github.com/pascalgaut/contributors" alt="Contributors">
        <img src="https://img.shields.io/github/contributors/mickael-kerjean/filestash" style="max-width:100%;">
    </a>
    <a href="https://hub.docker.com/r/machines/filestash" alt="Docker Hub">
        <img src="https://img.shields.io/docker/pulls/machines/filestash" style="max-width:100%;">
    </a>
    <a href="https://kiwiirc.com/nextclient/#irc://irc.libera.chat/#filestash?nick=guest??" alt="Chat on IRC">
        <img src="https://img.shields.io/badge/IRC-%23filestash-brightgreen.svg" style="max-width:100%;">
    </a>
</p>

<p align="center">
    A Dropbox-like file manager that let you manage your data anywhere it is located:<br>
    <a href="https://www.filestash.app/ftp-client.html">FTP</a> • FTPS • <a href="https://www.filestash.app/ssh-file-transfer.html">SFTP</a> • <a href="https://www.filestash.app/webdav-client.html">WebDAV</a> • Git • <a href="https://www.filestash.app/s3-browser.html">S3</a> • NFS • <a href="https://www.filestash.app/smb-client.html">SMB</a> • Artifactory • <a href="https://www.filestash.app/ldap-browser.html">LDAP</a> • Mysql <br>
       Storj • CardDAV • CalDAV • Backblaze B2 • <a href="https://www.filestash.app/s3-browser.html">Minio</a> <br>
               Dropbox • Google Drive
</p>
<p align="center">
    <a href="http://demo.filestash.app">
      <img src="https://raw.githubusercontent.com/mickael-kerjean/filestash_images/master/.assets/button_demo.png" alt="demo button" />
    </a>
</p>

# Features
- Manage your files from a browser
- Authentication middleware to connect to various source of user
- Flexible Share mechanism
- Chromecast support for images, music, and videos
- Video player
- Video transcoding (mov, mkv, avi, mpeg, and more)
- Image viewer
- Image transcoding (raw images from Nikon, Canon, and more)
- Photo management
- Audio player
- Shared links are full fledge network drive
- Office documents (docx, xlsx and more)
- Full org mode client ([documentation](https://www.filestash.app/2018/05/31/release-note-v0.1/))
- User friendly
- Mobile friendly
- Customisable
- Plugins
- Super fast
- Upload files and folders
- Download as zip
- Multiple cloud providers and protocols, easily extensible
- Nyan cat loader
- Emacs, VIM or Sublime keybindings `;)`
- Pluggable Search (default is recursive search)
- .. and many more

# Screenshots
<p align="center">
    <a href="https://demo.filestash.app">
        <img src="https://raw.githubusercontent.com/mickael-kerjean/filestash_images/master/.assets/navigation.gif" alt="user experience on navigation" />
    </a>
</p>
<p align="center">
    <a href="http://demo.filestash.app">
        <img src="https://raw.githubusercontent.com/mickael-kerjean/filestash_images/master/.assets/photo_management.gif" alt="user experience on medias" />
    </a>
</p>

# Core Principles and Ideas

Filestash started as an attempt to solve the "Dropbox problem" by abstracting the storage layer, allowing you to "bring your own backend" by implementing this interface:
```go
type IBackend interface {
	Init(params map[string]string, app *App) (IBackend, error) // constructor
	Ls(path string) ([]os.FileInfo, error)           // list files in a folder
	Cat(path string) (io.ReadCloser, error)          // download a file
	Mkdir(path string) error                         // create a folder
	Rm(path string) error                            // remove something
	Mv(from string, to string) error                 // rename something
	Save(path string, file io.Reader) error          // save a file
	Touch(path string) error                         // create a file
	LoginForm() Form                                 // dynamic form generation for the login
}
```

The software is effectively 2 distinct parts: Core and [Plugins](https://github.com/pascalgaut/filestash/tree/master/server/plugin). Plugins act as "Lego blocks" that shape all the opinions on how the file manager operates, such as who can do what and where (aka, the authorisation), how the search features work (e.g., recursive search, full text search, ...), how users are authenticated (e.g., passthrough to the storage layer, LDAP, OIDC, SAML, signed URLs), how the application behaves when opening different file types, etc...

The architecture is designed to enable implementers to build file managers that are fit for purpose. To cite an example, several control panels for elevators in Europe feature QR codes generated by Filestash. Technicians responsible for their maintenance scan these QR codes to access the complete history of the elevator, upload new records, access relevant external resources, and view a banner displaying location specific metadata from the ERP. What would have been a complex custom solution was made possible through the addition of two plugins: one for QR code generation and another for ERP integration. Everything else was off the shelf.

# Documentation
- [Getting started](https://www.filestash.app/docs/)
- [Installation](https://www.filestash.app/docs/install-and-upgrade/)
- [API](https://www.filestash.app/docs/api/)

# Support
- For companies -> [support contract](https://www.filestash.app/pricing/)
- For individuals -> [#filestash](https://kiwiirc.com/nextclient/#irc://irc.libera.chat/#filestash?nick=guest??) on IRC (libera.chat). To financially contribute to the project:
  - Bitcoin: `3LX5KGmSmHDj5EuXrmUvcg77EJxCxmdsgW`
  - [Open Collective](https://opencollective.com/filestash)

# Credits
- [Contributors](https://github.com/pascalgaut/filestash/graphs/contributors) and folks developing [awesome libraries](https://github.com/pascalgaut/filestash/blob/master/go.mod)
- This project is tested with BrowserStack
- Logo derived from the work of [ssnjrthegr8](https://github.com/ssnjrthegr8), Iconography from [flaticon](https://www.flaticon.com/), [fontawesome](https://fontawesome.com) and [material](https://material.io/icons/)
- [libvips](https://github.com/libvips/libvips) and [libraw](https://github.com/LibRaw/LibRaw). Those libraries are statically compiled in Filestash. Instructions to build Filestash is available [here](https://github.com/pascalgaut/filestash/blob/master/.drone.yml) and instructions to create your own static library for libvips and libraw is to be found [here](https://github.com/pascalgaut/filestash/tree/master/server/plugin/plg_image_light/deps)
