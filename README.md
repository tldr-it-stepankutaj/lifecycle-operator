# Operator Lifecycle Operator

**Operator Lifecycle Operator** je vlastní Kubernetes (OpenShift) operátor, který spravuje životní cyklus ostatních operátorů ve vašem clusteru. Umožňuje sledovat, detekovat nové verze a volitelně provádět automatický nebo manuální upgrade.

---

## ✨ Funkce

- 📦 Sledování `Subscription` pro určené operátory
- 🔍 Detekce nové verze pomocí `PackageManifest` z `CatalogSource`
- 🚀 Automatický upgrade (`autoUpgrade: true`)
- 📬 Webhook notifikace při nové verzi
- 📤 Podpora ručního upgradu pomocí CRD `OperatorUpgradeRequest`
- 🔧 CLI nástroj `oc-lifecycle` (Cobra-based)
- 🖥️ Možná integrace s OpenShift Console Pluginem (volitelné)

---

## 🚀 Rychlý start

### 1. Vytvoření projektu

```bash
operator-sdk init --domain tldr.cloud --repo github.com/tldr-it-stepankutaj/lifecycle-operator
operator-sdk create api --group lifecycle --version v1alpha1 --kind WatchedOperator --resource --controller
operator-sdk create api --group lifecycle --version v1alpha1 --kind OperatorUpgradeRequest --resource --controller
```

### 2. Instalace závislostí

```bash
go get github.com/operator-framework/api@v0.29.1
go get github.com/operator-framework/operator-lifecycle-manager@v0.26.0
go mod tidy
```

---

## 📄 CRD Příklady

### `WatchedOperator`

```yaml
apiVersion: lifecycle.tldr.cloud/v1alpha1
kind: WatchedOperator
metadata:
  name: prometheus-operator
spec:
  name: prometheus-operator
  namespace: openshift-monitoring
  channel: stable
  autoUpgrade: true
  webhookURL: https://hooks.slack.com/services/...
```

### `OperatorUpgradeRequest`

```yaml
apiVersion: lifecycle.tldr.cloud/v1alpha1
kind: OperatorUpgradeRequest
metadata:
  name: prometheus-upgrade
spec:
  operatorName: prometheus-operator
  namespace: openshift-monitoring
  targetCSV: prometheus-operator.v0.72.0
```

---

## 🔧 CLI (oc-lifecycle)

```bash
oc-lifecycle register prometheus-operator --channel stable
oc-lifecycle list
oc-lifecycle upgrade prometheus-operator --to v0.72.0
```

---

## 🧪 Vývoj a testování

```bash
make install
make run
```

---

## 📦 Build operátoru jako image

```bash
make docker-build docker-push IMG=quay.io/youruser/lifecycle-operator:latest
```

---

## 📬 Webhook JSON

```json
{
  "operator": "prometheus-operator",
  "current": "v0.71.0",
  "latest": "v0.72.0",
  "message": "New version available"
}
```

---

## 📃 Licence

MIT © 2024 [Stepan Kutaj](https://github.com/tldr-it-stepankutaj)
