# frozen_string_literal: true

class ProjectsController < ApplicationController
  before_action :authenticate_user!

  # GET /projects
  # GET /projects.json
  def index
    @projects = Project.all
  end

  # GET /projects/1
  # GET /projects/1.json
  def show
    set_project_with_flags
  end

  # GET /projects/new
  def new
    @project = current_user_projects.new
  end

  # GET /projects/1/edit
  def edit
    set_project
  end

  # POST /projects
  # POST /projects.json
  def create
    @project = current_user_projects.new(project_params)

    if @project.save
      redirect_to @project, notice: "Project was successfully created."
    else
      render :new
    end
  end

  # PATCH/PUT /projects/1
  # PATCH/PUT /projects/1.json
  def update
    set_project
    if @project.update(project_params)
      redirect_to @project, notice: "Project was successfully updated."
    else
      render :edit
    end
  end

  # DELETE /projects/1
  # DELETE /projects/1.json
  def destroy
    set_project
    @project.destroy
    redirect_to projects_url, notice: "Project was successfully destroyed."
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_project
      @project = current_user_projects.find(params[:id])
    end

    def set_project_with_flags
      @project = current_user_projects.includes(:flags).find(params[:id])
    end

    def current_user_projects
      current_user.projects
    end

    # Never trust parameters from the scary internet, only allow the white list through.
    def project_params
      params.require(:project).permit(:name)
    end
end
