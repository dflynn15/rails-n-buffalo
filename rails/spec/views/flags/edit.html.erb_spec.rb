require 'rails_helper'

RSpec.describe "flags/edit", type: :view do
  before(:each) do
    @flag = assign(:flag, Flag.create!(
      :name => "MyString",
      :enabled => false
    ))
  end

  it "renders the edit flag form" do
    render

    assert_select "form[action=?][method=?]", flag_path(@flag), "post" do

      assert_select "input[name=?]", "flag[name]"

      assert_select "input[name=?]", "flag[enabled]"
    end
  end
end
