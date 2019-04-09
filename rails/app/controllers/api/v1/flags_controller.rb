# frozen_string_literal: true

module Api::V1
  class FlagsController < ApiController
    def show
      render json: current_user.projects.find(params[:project_id]).flags.find(params[:id])
    end
  end
end
