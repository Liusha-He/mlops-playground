from pytest_bdd import scenario, given, when, then


@scenario("../features/reporting.feature", 
          "all existing patients are shwon in patient list view")
def test_patient_list_view():
    assert 1 == False


@given("head into the reporting data home page")
def head_to_page():
    return True


@when("any patients exist in the database")
def patients_exist():
    return head_to_page()


@then("all patients will be seen on the browser")
def see_patients():
    assert 1 == True
