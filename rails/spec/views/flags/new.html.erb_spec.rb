require 'rails_helper'

RSpec.describe "flags/new", type: :view do
  before(:each) do
    assign(:flag, Flag.new(
      :name => "MyString",
      :enabled => false
    ))
  end

  it "renders new flag form" do
    render

    assert_select "form[action=?][method=?]", flags_path, "post" do

      assert_select "input[name=?]", "flag[name]"

      assert_select "input[name=?]", "flag[enabled]"
    end
  end
end
