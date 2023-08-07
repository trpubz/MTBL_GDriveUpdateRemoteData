# MTBL_GDriveUpdateRemoteData
A utility app that locates local files from the Extract-Transform legs of the MTBL app for deposit into remote location.
The files are ingested via API for Loading leg of ETL app.

## Requirements
- Client Secret needs to be pulled from [OAuth](https://console.cloud.google.com/apis/credentials/oauthclient/) location with project already set up.
  - file needs to be stored in package directory called `credentials.json`
- First time `go run ` is performed, navigate to token exchange url provided in terminal and paste the returned token in the terminal.
- `FileIds` retrieved from Google Drive need to updated in the main along with updated file paths.