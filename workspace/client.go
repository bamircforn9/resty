package resty

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"
)

// ... (rest of client.go content with the fix applied) ...