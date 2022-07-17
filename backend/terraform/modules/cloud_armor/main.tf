resource "google_compute_security_policy" "policy" {
  name = var.name

  rule {
    action   = "deny(403)"
    priority = "200"
    match {
      expr {
        expression = "evaluatePreconfiguredExpr('sqli-stable', ['owasp-crs-v030001-id942420-sqli', 'owasp-crs-v030001-id942421-sqli', 'owasp-crs-v030001-id942431-sqli', 'owasp-crs-v030001-id942432-sqli','owasp-crs-v030001-id942460-sqli', 'owasp-crs-v030001-id942440-sqli', 'owasp-crs-v030001-id942450-sqli', 'owasp-crs-v030001-id942430-sqli','owasp-crs-v030001-id942200-sqli'])"
      }
    }
    description = "SQLインジェクション防御ルール"
  }

  rule {
    action   = "deny(403)"
    priority = "201"
    match {
      expr {
        expression = "evaluatePreconfiguredExpr('lfi-stable')"
      }
    }
    description = "ローカルファイルインクルード防御ルール"
  }

  rule {
    action   = "deny(403)"
    priority = "202"
    match {
      expr {
        expression = "evaluatePreconfiguredExpr('rfi-stable')"
      }
    }
    description = "リモートファイルインクルード防御ルール"
  }

  rule {
    action   = "deny(403)"
    priority = "203"
    match {
      expr {
        expression = "evaluatePreconfiguredExpr('rce-stable')"
      }
    }
    description = "リモートコード実行防御ルール"
  }

  rule {
    action   = "allow"
    priority = "2147483647"
    match {
      versioned_expr = "SRC_IPS_V1"
      config {
        src_ip_ranges = ["*"]
      }
    }
    description = "default rule"
  }
}