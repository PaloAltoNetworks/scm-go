# scm-go
SDK for interacting with Strata Cloud Manager.

[![GoDoc](https://godoc.org/github.com/PaloAltoNetworks/scm-go?status.svg)](https://godoc.org/github.com/PaloAltoNetworks/scm-go)

NOTE: This sdk code is auto-generated.

---
## Beta Release Disclaimer

**This software is a pre-release version and is not ready for production use.**

*   **No Warranty:** This software is provided "as is," without any warranty of any kind, either expressed or implied, including, but not limited to, the implied warranties of merchantability and fitness for a particular purpose.
*   **Instability:** The beta software may contain defects, may not operate correctly, and may be substantially modified or withdrawn at any time.
*   **Limitation of Liability:** In no event shall the authors or copyright holders be liable for any claim, damages, or other liability, whether in an action of contract, tort, or otherwise, arising from, out of, or in connection with the beta software or the use or other dealings in the beta software.
*   **Feedback:** We encourage and appreciate your feedback and bug reports. However, you acknowledge that any feedback you provide is non-confidential.

By using this software, you agree to these terms.
---


## Warranty
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

THIS SOFTWARE IS RELEASED AS A PROOF OF CONCEPT FOR EXPERIMENTAL PURPOSES ONLY. USE IT AT OWN RISK. THIS SOFTWARE IS NOT SUPPORTED.

## Using scm-go

In the project root scm-go, populate scm-config.json with the relevant parameters for auth_url, client_id, client_secret, host, protocol, scope etc.

```aiignore
{
  "auth_url": "",
  "client_id": "",
  "client_secret": "",
  "host": "",
  "logging": "quiet",
  "protocol": "https",
  "scope": "",
  "skip_verify_certificate": false
}
```

Then you can write a go program to test out the authentication.
There are tests provided in the tests directory for convenience.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	setup "github.com/paloaltonetworks/scm-go-v2"
	"github.com/paloaltonetworks/scm-go-v2/common"
	"github.com/paloaltonetworks/scm-go-v2/generated/network_services"
)

func main() {
	configPath := common.GetConfigPath()
	setupClient := &setup.Client{
		AuthFile:         configPath,
		CheckEnvironment: false,
	}

	fmt.Printf("Using config file: %s\n", setupClient.AuthFile)

	// Setup the client configuration
	err := setupClient.Setup()

	// Refresh JWT token
	ctx := context.Background()
	if setupClient.Jwt == "" {
		maxRetries := 3
		retryDelay := 2 * time.Second
		for i := 0; i < maxRetries; i++ {
			err = setupClient.RefreshJwt(ctx)
			if err == nil {
				break // Success, exit the loop
			}
			time.Sleep(retryDelay)
		}
		// Fail the test only after all retries have been exhausted.
	}

	// Create the network_services API client
	config := network_services.NewConfiguration()
	config.Host = setupClient.GetHost()
	config.Scheme = "https"

	// Create a custom HTTP client that includes the JWT token and logging
	if setupClient.HttpClient == nil {
		setupClient.HttpClient = &http.Client{}
	}

	// Wrap the transport with our logging transport
	if setupClient.HttpClient.Transport == nil {
		setupClient.HttpClient.Transport = http.DefaultTransport
	}
	setupClient.HttpClient.Transport = &common.LoggingRoundTripper{
		Wrapped: setupClient.HttpClient.Transport,
	}

	config.HTTPClient = setupClient.HttpClient

	// Set up the default header with JWT
	config.DefaultHeader = make(map[string]string)
	config.DefaultHeader["Authorization"] = "Bearer " + setupClient.Jwt
	config.DefaultHeader["x-auth-jwt"] = setupClient.Jwt

	apiClient := network_services.NewAPIClient(config)

	// Create a profile to retrieve.
	profileName := "test-ike-get-" + common.GenerateRandomString(6)
	profile := network_services.IkeCryptoProfiles{
		Folder:     common.StringPtr("Shared"),
		Name:       profileName,
		Hash:       []string{"sha512"},
		DhGroup:    []string{"group20"},
		Encryption: []string{"aes-128-gcm"},
	}

	reqCreate := apiClient.IKECryptoProfilesAPI.CreateIKECryptoProfiles(context.Background()).IkeCryptoProfiles(profile)
	createRes, _, err := reqCreate.Execute()
	log.Print(*createRes.Id)
}
```
