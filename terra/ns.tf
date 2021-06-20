resource "kubernetes_namespace" "ruslangr_dev" {
  metadata {
    annotations = {
      name = "dev ns"
    }

    labels = {
      mylabel = "dev"
    }

    name = "development"
  }
}

resource "kubernetes_namespace" "ruslangr_prod" {
  metadata {
    annotations = {
      name = "prod ns!!!"
    }

    labels = {
      mylabel = "prod"
    }

    name = "production"
  }
}

resource "kubernetes_namespace" "ruslangr_cicd" {
  metadata {
    annotations = {
      name = "cicd ns"
    }

    labels = {
      mylabel = "cicd"
    }

    name = "cicd"
  }
}


resource "kubernetes_namespace" "ruslangr_monitoring" {
  metadata {
    annotations = {
      name = "monit ns"
    }

    labels = {
      mylabel = "monitoring"
    }

    name = "monitoring"
  }
}