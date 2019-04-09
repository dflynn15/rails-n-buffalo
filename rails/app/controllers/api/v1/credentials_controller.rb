# frozen_string_literal: true

module Api::V1
  class CredentialsController < ApiController
    def me
      render json: current_user
    end
  end
end
