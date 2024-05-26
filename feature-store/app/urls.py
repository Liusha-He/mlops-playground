from django.conf.urls.static import static
from django.contrib import admin
from django.urls import path, include

from app import settings

urlpatterns = [
    path("", include("app.reporting.urls")),
    # path("quotes/", include("app.quotes.urls")),
    path("admin/", admin.site.urls),
] + static(settings.STATIC_URL, document_root=settings.STATIC_ROOT)
