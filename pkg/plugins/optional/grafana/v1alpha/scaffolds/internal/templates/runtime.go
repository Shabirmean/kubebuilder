/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package templates

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
)

var _ machinery.Template = &RuntimeManifest{}

// Kustomization scaffolds a file that defines the kustomization scheme for the prometheus folder
type RuntimeManifest struct {
	machinery.TemplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *RuntimeManifest) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("grafana", "controller-runtime-metrics.json")
	}

	f.TemplateBody = controllerRuntimeTemplate

	f.IfExistsAction = machinery.OverwriteFile

	return nil
}

// nolint: lll
const controllerRuntimeTemplate = `{
  "__inputs": [
    {
      "name": "DS_PROMETHEUS",
      "label": "Prometheus",
      "description": "",
      "type": "datasource",
      "pluginId": "prometheus",
      "pluginName": "Prometheus"
    }
  ],
  "__requires": [
    {
      "type": "datasource",
      "id": "prometheus",
      "name": "Prometheus",
      "version": "1.0.0"
    }
  ],
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": "-- Grafana --",
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "target": {
          "limit": 100,
          "matchAny": false,
          "tags": [],
          "type": "dashboard"
        },
        "type": "dashboard"
      }
    ]
  },
  "editable": true,
  "fiscalYearStartMonth": 0,
  "graphTooltip": 0,
  "links": [],
  "liveNow": false,
  "panels": [
    {
      "datasource": "${DS_PROMETHEUS}",
      "description": "Total number of reconciliation errors per controller",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "continuous-GrYlRd"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 20,
            "gradientMode": "scheme",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "smooth",
            "lineWidth": 3,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "cpm"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 5
      },
      "id": 6,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "targets": [
        {
          "datasource": "${DS_PROMETHEUS}",
          "editorMode": "code",
          "exemplar": true,
          "expr": "sum(rate(controller_runtime_reconcile_errors_total{job=\"$job\", namespace=\"$namespace\"}[5m])) by (instance, pod)",
          "interval": "",
          "legendFormat": "{{"{{instance}}"}} {{"{{pod}}"}}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "Reconciliation Error Count Per Controller",
      "type": "timeseries"
    },
    {
      "datasource": "${DS_PROMETHEUS}",
      "description": "Total number of reconciliation per controller",
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "continuous-GrYlRd"
          },
          "custom": {
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 20,
            "gradientMode": "scheme",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "smooth",
            "lineWidth": 3,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              },
              {
                "color": "red",
                "value": 80
              }
            ]
          },
          "unit": "cpm"
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 12,
        "y": 5
      },
      "id": 7,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom"
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "targets": [
        {
          "datasource": "${DS_PROMETHEUS}",
          "editorMode": "code",
          "exemplar": true,
          "expr": "sum(rate(controller_runtime_reconcile_total{job=\"$job\", namespace=\"$namespace\"}[5m])) by (instance, pod)",
          "interval": "",
          "legendFormat": "{{"{{instance}}"}} {{"{{pod}}"}}",
          "range": true,
          "refId": "A"
        }
      ],
      "title": "Reconciliation Total Count Per Controller",
      "type": "timeseries"
    }
  ],
  "refresh": "",
  "style": "dark",
  "tags": [],
  "templating": {
    "list": [
      {
        "datasource": "${DS_PROMETHEUS}",
        "definition": "label_values(controller_runtime_reconcile_total{namespace=~\"$namespace\"}, job)",
        "hide": 0,
        "includeAll": false,
        "multi": false,
        "name": "job",
        "options": [],
        "query": {
          "query": "label_values(controller_runtime_reconcile_total{namespace=~\"$namespace\"}, job)",
          "refId": "StandardVariableQuery"
        },
        "refresh": 2,
        "regex": "",
        "skipUrlSync": false,
        "sort": 0,
        "type": "query"
      },
      {
        "current": {
          "selected": false,
          "text": "observability",
          "value": "observability"
        },
        "datasource": "${DS_PROMETHEUS}",
        "definition": "label_values(controller_runtime_reconcile_total, namespace)",
        "hide": 0,
        "includeAll": false,
        "multi": false,
        "name": "namespace",
        "options": [],
        "query": {
          "query": "label_values(controller_runtime_reconcile_total, namespace)",
          "refId": "StandardVariableQuery"
        },
        "refresh": 1,
        "regex": "",
        "skipUrlSync": false,
        "sort": 0,
        "type": "query"
      },
      {
        "current": {
          "selected": false,
          "text": "All",
          "value": "$__all"
        },
        "datasource": "${DS_PROMETHEUS}",
        "definition": "label_values(controller_runtime_reconcile_total{namespace=~\"$namespace\", job=~\"$job\"}, pod)",
        "hide": 2,
        "includeAll": true,
        "label": "pod",
        "multi": true,
        "name": "pod",
        "options": [],
        "query": {
          "query": "label_values(controller_runtime_reconcile_total{namespace=~\"$namespace\", job=~\"$job\"}, pod)",
          "refId": "StandardVariableQuery"
        },
        "refresh": 2,
        "regex": "",
        "skipUrlSync": false,
        "sort": 0,
        "type": "query"
      }
    ]
  },
  "time": {
    "from": "now-15m",
    "to": "now"
  },
  "timepicker": {},
  "timezone": "",
  "title": "Controller-Runtime-Metrics",
  "weekStart": ""
}
`