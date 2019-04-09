# frozen_string_literal: true

module Api::V1
  class ProjectsController < ApiController
    def show
      render json: current_user.projects.find(params[:id])
    end
  end
end
