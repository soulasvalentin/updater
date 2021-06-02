Desktop multi-platform updater. Admins uploads their application to a public repository and end-users download missing or changed files only form the same repository

## How to use

### Parameters

- **remoteUrl**: Where the `manifest.json` and downloadable files are hosted. URL should be public.

#### With local config file

Create the following file named `updater.json`

```json
{
  "remoteUrl": "https://mysite.com/files"
}
```

#### Without local config file

Use the following execution arguments (will always be priority against the file settings)

- `-remoteUrl=https://mysite.com/files`

### Running the application

The applications supports the following execution commands

**build**
Generates the manifest file of the current directory (and children). The application itself is excluded from the manifest.

**sync**
Used by end users. Download remote manifest, evaluates which files are missing or changed and downloads originals.