package actions

import (
	"github.com/dflynn15/rails-n-buffalo/buffalo/flags/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Project)
// DB Table: Plural (projects)
// Resource: Plural (Projects)
// Path: Plural (/projects)
// View Template Folder: Plural (/templates/projects/)

// ProjectsResource is the resource for the Project model
type ProjectsResource struct {
	buffalo.Resource
}

// List gets all Projects. This function is mapped to the path
// GET /projects
func (v ProjectsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	user := c.Value("current_user").(*models.User)
	projects := &models.Projects{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Projects from the DB
	if err := q.Where("user_id = ?", user.ID).All(projects); err != nil {
		return errors.WithStack(err)
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, projects))
}

// Show gets the data for one Project. This function is mapped to
// the path GET /projects/{project_id}
func (v ProjectsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Project
	project := &models.Project{}

	// To find the Project the parameter project_id is used.
	if err := tx.Eager().Find(project, c.Param("project_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, project))
}

// New renders the form for creating a new Project.
// This function is mapped to the path GET /projects/new
func (v ProjectsResource) New(c buffalo.Context) error {
	return c.Render(200, r.Auto(c, &models.Project{}))
}

// Create adds a Project to the DB. This function is mapped to the
// path POST /projects
func (v ProjectsResource) Create(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	user := c.Value("current_user").(*models.User)

	// Allocate an empty Project
	project := &models.Project{}

	// Bind project to the html form elements
	if err := c.Bind(project); err != nil {
		return errors.WithStack(err)
	}

	project.UserID = user.ID

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(project)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, project))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "project.created.success"))
	// and redirect to the projects index page
	return c.Render(201, r.Auto(c, project))
}

// Edit renders a edit form for a Project. This function is
// mapped to the path GET /projects/{project_id}/edit
func (v ProjectsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Project
	project := &models.Project{}

	if err := tx.Find(project, c.Param("project_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, project))
}

// Update changes a Project in the DB. This function is mapped to
// the path PUT /projects/{project_id}
func (v ProjectsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Project
	project := &models.Project{}

	if err := tx.Find(project, c.Param("project_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Project to the html form elements
	if err := c.Bind(project); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(project)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, project))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", T.Translate(c, "project.updated.success"))
	// and redirect to the projects index page
	return c.Render(200, r.Auto(c, project))
}

// Destroy deletes a Project from the DB. This function is mapped
// to the path DELETE /projects/{project_id}
func (v ProjectsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Allocate an empty Project
	project := &models.Project{}

	// To find the Project the parameter project_id is used.
	if err := tx.Eager().Find(project, c.Param("project_id")); err != nil {
		return c.Error(404, err)
	}

	// Destroy all of the Flags
	if err := tx.Destroy(project.Flags); err!=nil {
		return errors.WithStack(err)
	}

	if err := tx.Destroy(project); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", T.Translate(c, "project.destroyed.success"))
	// Redirect to the projects index page
	return c.Render(200, r.Auto(c, project))
}
