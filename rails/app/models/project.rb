class Project < ApplicationRecord
  has_many :flags, dependent: :destroy
  validates_presence_of :name
end
