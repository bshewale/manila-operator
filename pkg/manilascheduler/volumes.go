package manilascheduler

import (
	manilav1 "github.com/openstack-k8s-operators/manila-operator/api/v1beta1"
	"github.com/openstack-k8s-operators/manila-operator/pkg/manila"
	corev1 "k8s.io/api/core/v1"
)

// GetVolumes -
func GetVolumes(parentName string, name string, extraVol []manilav1.ManilaExtraVolMounts) []corev1.Volume {
	var config0640AccessMode int32 = 0640

	schedulerVolumes := []corev1.Volume{
		{
			Name: "config-data-custom",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					DefaultMode: &config0640AccessMode,
					LocalObjectReference: corev1.LocalObjectReference{
						Name: name + "-config-data",
					},
				},
			},
		},
	}

	return append(manila.GetVolumes(parentName, extraVol, manila.ManilaSchedulerPropagation), schedulerVolumes...)
}

// GetInitVolumeMounts - ManilaScheduler init task VolumeMounts
func GetInitVolumeMounts(extraVol []manilav1.ManilaExtraVolMounts) []corev1.VolumeMount {

	customConfVolumeMount := corev1.VolumeMount{
		Name:      "config-data-custom",
		MountPath: "/var/lib/config-data/custom",
		ReadOnly:  true,
	}

	return append(manila.GetInitVolumeMounts(extraVol, manila.ManilaSchedulerPropagation), customConfVolumeMount)
}

// GetVolumeMounts - ManilaScheduler VolumeMounts
func GetVolumeMounts(extraVol []manilav1.ManilaExtraVolMounts) []corev1.VolumeMount {
	return manila.GetVolumeMounts(extraVol, manila.ManilaSchedulerPropagation)
}
