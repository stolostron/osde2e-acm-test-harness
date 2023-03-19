package tests

import (
	"log"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/openshift/osde2e-acm-test-harness/pkg/metadata"

	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"k8s.io/client-go/rest"
)

// Using Reference Addon for test
// var crdName = "referenceaddons.reference.addons.managed.openshift.io"
var crdACMName = "multiclusterhubs.operator.open-cluster-management.io"

var _ = ginkgo.Describe("ACM osd addon", func() {
	defer ginkgo.GinkgoRecover()
	config, err := rest.InClusterConfig()

	if err != nil {
		panic(err)
	}

	ginkgo.It("should have crd "+crdACMName+".", func() {
		apiextensions, err := clientset.NewForConfig(config)
		Expect(err).NotTo(HaveOccurred())

		// Make sure the CRD exists
		result, err := apiextensions.ApiextensionsV1().CustomResourceDefinitions().Get(crdACMName, v1.GetOptions{})

		if err != nil {
			log.Printf("CRD not found: %v", err.Error())
			metadata.Instance.FoundCRD = false
		} else {
			log.Printf("CRD found: %v", result)
			metadata.Instance.FoundCRD = true
		}

		Expect(err).NotTo(HaveOccurred())
	}, float64(30))

	ginkgo.It("has a default passthrough secret", func() {
		k8s, err := kubernetes.NewForConfig(config)
		Expect(err).NotTo(HaveOccurred())

		sec, err := k8s.CoreV1().Secrets("osde2e-ci-secrets").Get("ci-secrets", v1.GetOptions{})
		// Expect(sec.Data["testkey"]).ToNot(BeNil())
		Expect(sec.Name).NotTo(BeNil())
		Expect(err).NotTo(HaveOccurred())

	}, float64(30))

	ginkgo.It("can deploy the ACM instance manifest", func() {

	})

	ginkgo.It("can verify the mch operator is succesful", func() {

	})

	ginkgo.It("can verify the mce operator is succesful", func() {

	})

	ginkgo.It("can enable the hypershift addon on the local-cluster", func() {

	})

	ginkgo.It("can run the hypershift addon controller without errors", func() {

	})

	ginkgo.It("can delete the passthrough secret", func() {

	})

})
