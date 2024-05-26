import pytest
from django.db import connection
from testcontainers.postgres import PostgresContainer


@pytest.fixture(scope="session")
def postgres_container(request):
    # Create a PostgreSQL container
    container = PostgresContainer("postgres:latest")
    container.start()     
    
    request.addfinalizer(lambda: container.stop())

    return container


@pytest.fixture
def setup_database(postgres_container):
    connection_params = postgres_container.get_connection_params()
    settings_override = {
        "DATABASES": {
            "default": {
                "ENGINE": "django.db.backends.postgresql",
                "NAME": connection_params['database'],
                "USER": connection_params['user'],
                "PASSWORD": connection_params['password'],
                "HOST": connection_params['host'],
                "PORT": connection_params['port'],
            }
        }
    }
    
    with connection.settings(**settings_override):
        # load init data here
        yield
