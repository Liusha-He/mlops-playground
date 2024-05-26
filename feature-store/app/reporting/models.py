from django.db import models


class Patient(models.Model):
    lastName = models.CharField(max_length=50)
    firstName = models.CharField(max_length=50)
    age = models.IntegerField()


class ClinicalData(models.Model):
    COMPONENT_NAMES = [("hw", "Height/Weight"), ("bp", "Blood Pressure"), ("heartrate", "Heart Rate")]

    componentName = models.CharField(
        max_length=20,
        choices=COMPONENT_NAMES,
    )
    componentValue = models.CharField(max_length=20)
    measuredDateTime = models.DateTimeField(auto_now_add=True)
    patient = models.ForeignKey(Patient, on_delete=models.CASCADE)
