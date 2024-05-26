from django.http import HttpRequest, HttpResponse
from django.shortcuts import render, redirect
from django.urls import reverse_lazy
from django.views.generic import ListView, CreateView, UpdateView, DeleteView

from app.reporting.forms import ClinicalDataForm
from app.reporting.models import Patient

class PatientListView(ListView):
    model = Patient


class PatientCreateView(CreateView):
    template_name = "reporting/patient_form.html"
    model = Patient
    success_url = reverse_lazy("patients")
    fields = ["firstName", "lastName", "age"]


class PatientUpdateView(UpdateView):
    template_name = "reporting/patient_form.html"
    model = Patient
    success_url = reverse_lazy("patients")
    fields = ["firstName", "lastName", "age"]


class PatientDeleteView(DeleteView):
    template_name = "reporting/patient_confirm_delete.html"
    model = Patient
    success_url = reverse_lazy("patients")


def add_clinical_data(req: HttpRequest, **kwargs) -> HttpResponse:
    form = ClinicalDataForm()
    patient = Patient.objects.get(id=kwargs["pk"])
    
    if req.method == "POST":
        form = ClinicalDataForm(req.POST)
        if form.is_valid():
            form.save()
        return redirect("/")
    return render(req, 
                  "reporting/clinical_data_form.html",
                  {
                      "form": form,
                      "patient": patient,
                  })
