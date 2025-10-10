<!--
{
	"nav_order": 2
}
-->

# Quick Start

here we can work on thing that we can do more.

### Init Doc Workstation

Its very simple to create a Doc Workstation to work.
this ask you to pick a name for you docs.

```batch
mdlook init .
```

### Watch the Doc

then you can watch your doc. open the brwoser the default host is `localhost:8080`

```batch
mdlook watch .
```

### Export & Bundle

after you finished your work and its ready to make that public do this command.
this bundle you work to be as servable things.

```batch
mdlook export --output <path>
```

> By default it export to `dist` folder.

### Serve (Production)

then you can serve your bundled documentation for production use. This will start a server to host your exported docs.

```batch
mdlook serve <path>
```
