package response

import (
	"fmt"
	"net/http"
	"runtime"

	pkgErrors "gitlab.com/tantai-smap/authenticate-api/pkg/errors"
	"gitlab.com/tantai-smap/authenticate-api/pkg/telegram"

	"github.com/gin-gonic/gin"
)

// Resp is the response format.
type Resp struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	Errors    any    `json:"errors,omitempty"`
}

// NewOKResp returns a new OK response with the given data.
func NewOKResp(data any) Resp {
	return Resp{
		ErrorCode: 0,
		Message:   "Success",
		Data:      data,
	}
}

// Ok returns a new OK response with the given data.
func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, NewOKResp(data))
}

// Unauthorized returns a new Unauthorized response with the given data.
func Unauthorized(c *gin.Context) {
	c.JSON(parseError(pkgErrors.NewUnauthorizedHTTPError(), c, nil, nil))
}

func Forbidden(c *gin.Context) {
	c.JSON(parseError(pkgErrors.NewForbiddenHTTPError(), c, nil, nil))
}

func parseError(err error, c *gin.Context, t *telegram.TeleBot, chatID *int64) (int, Resp) {
	//print error . type
	switch parsedErr := err.(type) {
	case *pkgErrors.ValidationError:
		return http.StatusBadRequest, Resp{
			ErrorCode: parsedErr.Code,
			Message:   parsedErr.Error(),
		}
	case *pkgErrors.PermissionError:
		return http.StatusBadRequest, Resp{
			ErrorCode: parsedErr.Code,
			Message:   parsedErr.Error(),
		}
	case *pkgErrors.PaymentError:
		return http.StatusForbidden, Resp{
			ErrorCode: parsedErr.Code,
			Message:   parsedErr.Error(),
		}
	case *pkgErrors.ValidationErrorCollector:
		return http.StatusBadRequest, Resp{
			ErrorCode: ValidationErrorCode,
			Message:   ValidationErrorMsg,
			Errors:    parsedErr.Errors(),
		}
	case *pkgErrors.PermissionErrorCollector:
		return http.StatusBadRequest, Resp{
			ErrorCode: PermissionErrorCode,
			Message:   PermissionErrorMsg,
			Errors:    parsedErr.Errors(),
		}
	case *pkgErrors.PaymentErrorCollector:
		return http.StatusBadRequest, Resp{
			ErrorCode: PaymentErrorCode,
			Message:   PaymentErrorMsg,
			Errors:    parsedErr.Errors(),
		}
	case *pkgErrors.HTTPError:
		statusCode := parsedErr.StatusCode
		if statusCode == 0 {
			statusCode = http.StatusBadRequest
		}

		return statusCode, Resp{
			ErrorCode: parsedErr.Code,
			Message:   parsedErr.Message,
		}
	default:
		if t != nil && chatID != nil {
			panic(err)
		}

		stackTrace := captureStackTrace()
		sendServerTelegramMessageAsync(buildInternalServerErrorDataForReportBug(err.Error(), stackTrace, c), *t, *chatID)

		return http.StatusInternalServerError, Resp{
			ErrorCode: 500,
			Message:   DefaultErrorMessage,
		}
	}
}

// Error returns a new Error response with the given error.
func Error(c *gin.Context, err error) {
	c.JSON(parseError(err, c, nil, nil))
}

// HttpError returns a new Error response with the given error.
func HttpError(c *gin.Context, err *pkgErrors.HTTPError) {
	c.JSON(parseError(err, c, nil, nil))
}

// ErrorMapping is a map of error to HTTPError.
type ErrorMapping map[error]*pkgErrors.HTTPError

// ErrorWithMap returns a new Error response with the given error.
func ErrorWithMap(c *gin.Context, err error, eMap ErrorMapping) {
	if httpErr, ok := eMap[err]; ok {
		Error(c, httpErr)
		return
	}

	Error(c, err)
}

func PanicError(c *gin.Context, err any, t telegram.TeleBot, tIDs int64) {
	if err == nil {
		c.JSON(parseError(nil, c, &t, &tIDs))
	} else {
		c.JSON(parseError(err.(error), c, &t, &tIDs))
	}
}

func captureStackTrace() []string {
	var pcs [defaultStackTraceDepth]uintptr
	n := runtime.Callers(2, pcs[:])
	if n == 0 {
		return nil
	}

	var stackTrace []string
	for _, pc := range pcs[:n] {
		f := runtime.FuncForPC(pc)
		if f != nil {
			file, line := f.FileLine(pc)
			stackTrace = append(stackTrace, fmt.Sprintf("%s:%d %s", file, line, f.Name()))
		}
	}

	return stackTrace
}
