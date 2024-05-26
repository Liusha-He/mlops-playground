from django.urls import path

from app.reporting import views


urlpatterns = [
    path("", views.PatientListView.as_view(), name="patients"),
    path("create", views.PatientCreateView.as_view(), name="create_patient"),
    path("update/<int:pk>", views.PatientUpdateView.as_view(), name="update_patient"),
    path("delete/<int:pk>", views.PatientDeleteView.as_view(), name="delete_patient"),
    path("addData/<int:pk>", views.add_clinical_data, name="add_clinical_data"),
]
