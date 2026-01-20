package cloud

import (
	"context"
	"fmt"
	"iter"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/capcom6/go-restkit"
	"github.com/softc24/evotor-go/helpers"
)

type Client struct {
	*restkit.Client
}

func NewClient(cfg ClientConfig) (*Client, error) {
	if cfg.BaseURL == "" {
		cfg.BaseURL = DefaultURL
	}

	rest, err := restkit.NewClient(restkit.Config{
		Client:  cfg.Client,
		BaseURL: cfg.BaseURL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return &Client{
		Client: rest,
	}, nil
}

// GetDevices returns a sequence of devices and a function to check for errors.
// The sequence is lazily loaded from the server using the provided token.
// The server will return a paginated list of devices, and the sequence will
// yield each device in the list. If there are more devices available, the
// sequence will automatically request the next page from the server.
// If an error occurs while loading the devices, the error function will
// return the error.
func (c *Client) GetDevices(ctx context.Context, token string) (iter.Seq[Device], func() error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	params := url.Values{}

	return helpers.GetPagedData[Device](
		ctx,
		c.Client,
		"/devices",
		params,
		headers,
	)
}

// GetStores returns a sequence of stores and a function to check for errors.
// The sequence is lazily loaded from the server using the provided token.
// The server will return a paginated list of stores, and the sequence will
// yield each store in the list. If there are more stores available, the
// sequence will automatically request the next page from the server.
// If an error occurs while loading the stores, the error function will
// return the error.
func (c *Client) GetStores(ctx context.Context, token string) (iter.Seq[Store], func() error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	params := url.Values{}

	return helpers.GetPagedData[Store](
		ctx,
		c.Client,
		"/stores",
		params,
		headers,
	)
}

// GetEmployees returns a sequence of employees and a function to check for errors.
// The sequence is lazily loaded from the server using the provided token.
// The server will return a paginated list of employees, and the sequence will
// yield each employee in the list. If there are more employees available, the
// sequence will automatically request the next page from the server.
// If an error occurs while loading the employees, the error function will
// return the error.
func (c *Client) GetEmployees(ctx context.Context, token string) (iter.Seq[Employee], func() error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	params := url.Values{}

	return helpers.GetPagedData[Employee](
		ctx,
		c.Client,
		"/employees",
		params,
		headers,
	)
}

// GetEmployee returns a specific employee by ID.
// The method makes a GET request to /employees/{employee-id} and returns the Employee object.
// Returns the Employee and an error if the request fails or the employee is not found.
func (c *Client) GetEmployee(ctx context.Context, token string, employeeID string) (*Employee, error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}

	var employee Employee
	path := fmt.Sprintf("/employees/%s", url.PathEscape(employeeID))

	err := c.Client.Do(ctx, http.MethodGet, path, headers, nil, &employee)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee %s: %w", employeeID, err)
	}

	return &employee, nil
}

// GetRoles returns a sequence of roles and a function to check for errors.
// The sequence is lazily loaded from the server using the provided token.
// The server will return a paginated list of roles, and the sequence will
// yield each role in the list. If there are more roles available, the
// sequence will automatically request the next page from the server.
// If an error occurs while loading the roles, the error function will
// return the error.
func (c *Client) GetRoles(ctx context.Context, token string) (iter.Seq[Role], func() error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	params := url.Values{}

	return helpers.GetPagedData[Role](
		ctx,
		c.Client,
		"/employees/roles",
		params,
		headers,
	)
}

// GetDocuments returns a sequence of documents and a function to check for errors.
// The sequence is lazily loaded from the server using the provided token.
// The server will return a paginated list of documents, and the sequence will
// yield each document in the list. If there are more documents available, the
// sequence will automatically request the next page from the server.
// If an error occurs while loading the documents, the error function will
// return the error.
//
// Parameters:
//   - storeID: Store identifier (required)
//   - since: Documents created from this date (in milliseconds, optional - use 0 to skip)
//   - until: Documents created before this date (in milliseconds, optional - use 0 to skip)
//   - docTypes: Filter by document types (optional - empty string to skip)
//
// Note: since, until, and docTypes should only be specified in the first request.
func (c *Client) GetDocuments(
	ctx context.Context,
	token string,
	storeID string,
	since int64,
	until int64,
	docTypes string,
) (iter.Seq[Document], func() error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	params := url.Values{}

	// Add optional query parameters only if provided
	if since > 0 {
		params.Set("since", strconv.FormatInt(since, 10))
	}
	if until > 0 {
		params.Set("until", strconv.FormatInt(until, 10))
	}
	if docTypes != "" {
		params.Set("type", docTypes)
	}

	path := fmt.Sprintf("/stores/%s/documents", url.PathEscape(storeID))

	return helpers.GetPagedData[Document](
		ctx,
		c.Client,
		path,
		params,
		headers,
	)
}

// GetDocument retrieves a document by ID.
// Parameters:
//   - token: Authorization token in Bearer format.
//   - storeID: Store identifier.
//   - documentID: Document identifier.
//
// Returns:
//   - *Document: Pointer to the retrieved document.
//   - error: Error if the request fails or document is not found.
func (c *Client) GetDocument(ctx context.Context, token string, storeID, documentID string) (*Document, error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	var doc Document
	path := fmt.Sprintf("/stores/%s/documents/%s", url.PathEscape(storeID), url.PathEscape(documentID))
	err := c.Client.Do(ctx, http.MethodGet, path, headers, nil, &doc)
	if err != nil {
		return nil, fmt.Errorf("failed to get document %s: %w", documentID, err)
	}
	return &doc, nil
}

// GetProductGroups returns a sequence of product groups and a function to check for errors.
// The sequence is lazily loaded from the server using the provided token.
// The server will return a paginated list of product groups, and the sequence will
// yield each product group in the list. If there are more product groups available, the
// sequence will automatically request the next page from the server.
// If an error occurs while loading the product groups, the error function will
// return the error.
//
// Parameters:
//   - storeID: Store identifier (required)
//   - since: Product groups updated from this date (in milliseconds, optional - use 0 to skip)
//   - isRemoved: If true, returns only removed product groups (optional)
//   - ids: Product group identifiers (optional)
//
// Note: since, isRemoved, and ids should only be specified in the first request for pagination.
func (c *Client) GetProductGroups(
	ctx context.Context,
	token string,
	storeID string,
	since int64,
	isRemoved bool,
	ids []string,
) (iter.Seq[ProductGroup], func() error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	params := url.Values{}

	// Add optional query parameters only if provided
	if since > 0 {
		params.Set("since", strconv.FormatInt(since, 10))
	}
	if isRemoved {
		params.Set("isRemoved", "true")
	}
	if len(ids) > 0 {
		params.Set("id", strings.Join(ids, ","))
	}

	path := fmt.Sprintf("/stores/%s/product-groups", url.PathEscape(storeID))

	return helpers.GetPagedData[ProductGroup](
		ctx,
		c.Client,
		path,
		params,
		headers,
	)
}

// GetProductGroup retrieves a product group by ID.
// Parameters:
//   - token: Authorization token in Bearer format.
//   - storeID: Store identifier.
//   - productGroupID: Product group identifier.
//
// Returns:
//   - *ProductGroup: Pointer to the retrieved product group.
//   - error: Error if the request fails or product group is not found.
func (c *Client) GetProductGroup(
	ctx context.Context,
	token string,
	storeID, productGroupID string,
) (*ProductGroup, error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	var productGroup ProductGroup
	path := fmt.Sprintf("/stores/%s/product-groups/%s", url.PathEscape(storeID), url.PathEscape(productGroupID))
	err := c.Client.Do(ctx, http.MethodGet, path, headers, nil, &productGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to get product group %s: %w", productGroupID, err)
	}
	return &productGroup, nil
}

// GetProducts returns a sequence of products and a function to check for errors.
// The sequence is lazily loaded from the server using the provided token.
// The server will return a paginated list of products, and the sequence will
// yield each product in the list. If there are more products available, the
// sequence will automatically request the next page from the server.
// If an error occurs while loading the products, the error function will
// return the error.
//
// Parameters:
//   - storeID: Store identifier (required)
//   - since: Products updated from this date (in milliseconds, optional - use 0 to skip)
//   - isRemoved: If true, returns only removed products (optional)
//   - ids: Product identifiers (optional)
//   - fields: Fields to include in the response (optional)
//
// Note: since, isRemoved, ids, and fields should only be specified in the first request for pagination.
func (c *Client) GetProducts(
	ctx context.Context,
	token string,
	storeID string,
	since int64,
	isRemoved bool,
	ids []string,
	fields []string,
) (iter.Seq[Product], func() error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	params := url.Values{}

	// Add optional query parameters only if provided
	if since > 0 {
		params.Set("since", strconv.FormatInt(since, 10))
	}
	if isRemoved {
		params.Set("isRemoved", "true")
	}
	if len(ids) > 0 {
		params.Set("id", strings.Join(ids, ","))
	}
	if len(fields) > 0 {
		params.Set("fields", strings.Join(fields, ","))
	}

	path := fmt.Sprintf("/stores/%s/products", url.PathEscape(storeID))

	return helpers.GetPagedData[Product](
		ctx,
		c.Client,
		path,
		params,
		headers,
	)
}

// GetProduct retrieves a product by ID.
// Parameters:
//   - token: Authorization token in Bearer format.
//   - storeID: Store identifier.
//   - productID: Product identifier.
//
// Returns:
//   - *Product: Pointer to the retrieved product.
//   - error: Error if the request fails or product is not found.
func (c *Client) GetProduct(ctx context.Context, token string, storeID, productID string) (*Product, error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	var product Product
	path := fmt.Sprintf("/stores/%s/products/%s", url.PathEscape(storeID), url.PathEscape(productID))
	err := c.Client.Do(ctx, http.MethodGet, path, headers, nil, &product)
	if err != nil {
		return nil, fmt.Errorf("failed to get product %s: %w", productID, err)
	}
	return &product, nil
}

// GetBulkTasks returns a sequence of bulk tasks and a function to check for errors.
// The sequence is lazily loaded from the server using the provided token.
// The server will return a paginated list of bulk tasks, and the sequence will
// yield each bulk task in the list. If there are more bulk tasks available, the
// sequence will automatically request the next page from the server.
// If an error occurs while loading the bulk tasks, the error function will
// return the error.
//
// Parameters:
//   - types: Comma-separated list of task types (required, e.g. "product,product-group")
//   - statuses: Comma-separated list of task statuses (optional, e.g. "RUNNING,COMPLETED")
//
// Note: types, and statuses should only be specified in the first request for pagination.
func (c *Client) GetBulkTasks(
	ctx context.Context,
	token string,
	types string,
	statuses []string,
) (iter.Seq[BulkTask], func() error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	params := url.Values{}

	// type is required
	params.Set("type", types)

	// Add optional query parameters only if provided
	if len(statuses) > 0 {
		params.Set("status", strings.Join(statuses, ","))
	}

	return helpers.GetPagedData[BulkTask](
		ctx,
		c.Client,
		"/bulks",
		params,
		headers,
	)
}

// GetBulkTask retrieves detailed information about a bulk task by ID.
// Parameters:
//   - token: Authorization token in Bearer format.
//   - bulkID: Bulk task identifier.
//
// Returns:
//   - *BulkTask: Pointer to the retrieved bulk task with details.
//   - error: Error if the request fails or task is not found.
func (c *Client) GetBulkTask(ctx context.Context, token string, bulkID string) (*BulkTask, error) {
	headers := http.Header{
		"Authorization": []string{"Bearer " + token},
		"Accept":        []string{"application/json"},
	}
	var bulkTask BulkTask
	path := fmt.Sprintf("/bulks/%s", url.PathEscape(bulkID))
	err := c.Client.Do(ctx, http.MethodGet, path, headers, nil, &bulkTask)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk task %s: %w", bulkID, err)
	}
	return &bulkTask, nil
}
