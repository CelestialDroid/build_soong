// Copyright 2023 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package genrule

var (
	DepfileAllowList = []string{
		// go/keep-sorted start
		"depfile_allowed_for_test",
		"gen_uwb_core_proto",
		"libtextclassifier_fbgen_actions_actions-entity-data",
		"libtextclassifier_fbgen_actions_actions_model",
		"libtextclassifier_fbgen_annotator_datetime_datetime",
		"libtextclassifier_fbgen_annotator_entity-data",
		"libtextclassifier_fbgen_annotator_experimental_experimental",
		"libtextclassifier_fbgen_annotator_model",
		"libtextclassifier_fbgen_annotator_person_name_person_name_model",
		"libtextclassifier_fbgen_lang_id_common_flatbuffers_embedding-network",
		"libtextclassifier_fbgen_lang_id_common_flatbuffers_model",
		"libtextclassifier_fbgen_utils_codepoint-range",
		"libtextclassifier_fbgen_utils_container_bit-vector",
		"libtextclassifier_fbgen_utils_flatbuffers_flatbuffers",
		"libtextclassifier_fbgen_utils_flatbuffers_flatbuffers_test",
		"libtextclassifier_fbgen_utils_grammar_rules",
		"libtextclassifier_fbgen_utils_grammar_semantics_expression",
		"libtextclassifier_fbgen_utils_grammar_testing_value",
		"libtextclassifier_fbgen_utils_i18n_language-tag",
		"libtextclassifier_fbgen_utils_intents_intent-config",
		"libtextclassifier_fbgen_utils_lua_utils_tests",
		"libtextclassifier_fbgen_utils_normalization",
		"libtextclassifier_fbgen_utils_resources",
		"libtextclassifier_fbgen_utils_tflite_text_encoder_config",
		"libtextclassifier_fbgen_utils_tokenizer",
		"libtextclassifier_fbgen_utils_zlib_buffer",
		"tflite_support_metadata_schema",
		"tflite_support_spm_config",
		"tflite_support_spm_encoder_config",
		// go/keep-sorted end
	}

	SandboxingDenyModuleList = []string{
		// go/keep-sorted start
		"CompilationTestCases_package-dex-usage",
		"ControlEnvProxyServerProto_cc",
		"ControlEnvProxyServerProto_h",
		"CtsApkVerityTestDebugFiles",
		"FrontendStub_cc",
		"FrontendStub_h",
		"HeadlessBuildTimestamp",
		"ImageProcessing-rscript",
		"ImageProcessing2-rscript",
		"ImageProcessingJB-rscript",
		"MultiDexLegacyTestApp_genrule",
		"PackageManagerServiceServerTests_apks_as_resources",
		"PacketStreamerStub_cc",
		"PacketStreamerStub_h",
		"RSTest-rscript",
		"RSTest_v11-rscript",
		"RSTest_v14-rscript",
		"RSTest_v16-rscript",
		"Refocus-rscript",
		"RsBalls-rscript",
		"ScriptGroupTest-rscript",
		"TracingVMProtoStub_cc",
		"TracingVMProtoStub_h",
		"UpdatableSystemFontTest_NotoColorEmojiV0.sig",
		"UpdatableSystemFontTest_NotoColorEmojiV0.ttf",
		"UpdatableSystemFontTest_NotoColorEmojiVPlus1.sig",
		"UpdatableSystemFontTest_NotoColorEmojiVPlus1.ttf",
		"UpdatableSystemFontTest_NotoColorEmojiVPlus2.sig",
		"UpdatableSystemFontTest_NotoColorEmojiVPlus2.ttf",
		"VehicleServerProtoStub_cc",
		"VehicleServerProtoStub_cc@2.0-grpc-trout",
		"VehicleServerProtoStub_cc@default-grpc",
		"VehicleServerProtoStub_h",
		"VehicleServerProtoStub_h@2.0-grpc-trout",
		"VehicleServerProtoStub_h@default-grpc",
		"aidl-golden-test-build-hook-gen",
		"aidl_camera_build_version",
		"android-cts-verifier",
		"android-support-multidex-instrumentation-version",
		"android-support-multidex-version",
		"angle_commit_id",
		"apexer_test_host_tools",
		"atest_integration_fake_src",
		"authfs_test_apk_assets",
		"awkgram.tab.h",
		"bluetooth_core_rust_packets",
		"c2hal_test_genc++",
		"c2hal_test_genc++_headers",
		"camera-its",
		"checkIn-service-stub-lite",
		"chre_atoms_log.h",
		"com.android.apex.test.bar_stripped",
		"com.android.apex.test.baz_stripped",
		"com.android.apex.test.foo_stripped",
		"com.android.apex.test.pony_stripped",
		"com.android.apex.test.sharedlibs_generated",
		"com.android.apex.test.sharedlibs_secondary_generated",
		"common-profile-text-protos",
		"core-tests-smali-dex",
		"cronet_aml_base_android_runtime_jni_headers",
		"cronet_aml_base_android_runtime_jni_headers__testing",
		"cronet_aml_base_android_runtime_unchecked_jni_headers",
		"cronet_aml_base_android_runtime_unchecked_jni_headers__testing",
		"deqp_spvtools_update_build_version",
		"egl_extensions_functions_hdr",
		"egl_functions_hdr",
		"emp_ematch.yacc.c",
		"emp_ematch.yacc.h",
		"fdt_test_tree_empty_memory_range_dtb",
		"fdt_test_tree_multiple_memory_ranges_dtb",
		"fdt_test_tree_one_memory_range_dtb",
		"futility_cmds",
		"gd_hci_packets_python3_gen",
		"gd_smp_packets_python3_gen",
		"gen_corrupt_rebootless_apex",
		"gen_corrupt_superblock_apex",
		"gen_key_mismatch_capex",
		"gen_manifest_mismatch_apex_no_hashtree",
		"generate_hash_v1",
		"gles1_core_functions_hdr",
		"gles1_extensions_functions_hdr",
		"gles2_core_functions_hdr",
		"gles2_extensions_functions_hdr",
		"gles31_only_functions_hdr",
		"gles3_only_functions_hdr",
		"hci_packets_python3_gen",
		"hidl2aidl_test_gen_aidl",
		"hidl2aidl_translate_cpp_test_gen_headers",
		"hidl2aidl_translate_cpp_test_gen_src",
		"hidl2aidl_translate_java_test_gen_src",
		"hidl2aidl_translate_ndk_test_gen_headers",
		"hidl2aidl_translate_ndk_test_gen_src",
		"hidl_cpp_impl_test_gen-headers",
		"hidl_cpp_impl_test_gen-sources",
		"hidl_error_test_gen",
		"hidl_export_test_gen-headers",
		"hidl_format_test_diff",
		"hidl_hash_test_gen",
		"hidl_hash_version_gen",
		"hidl_java_impl_test_gen",
		"lib-test-profile-text-protos",
		"libbssl_sys_src_nostd",
		"libc_musl_sysroot_bits",
		"libchrome-crypto-include",
		"libchrome-include",
		"libcore-non-cts-tests-txt",
		"libmojo_jni_headers",
		"libxml2_schema_fuzz_corpus",
		"libxml2_xml_fuzz_corpus",
		"link_layer_packets_python3_gen",
		"llcp_packets_python3_gen",
		"ltp_config_arm",
		"ltp_config_arm_64",
		"ltp_config_arm_64_hwasan",
		"ltp_config_arm_64_lowmem",
		"ltp_config_arm_64_lowmem_hwasan",
		"ltp_config_arm_lowmem",
		"ltp_config_riscv_64",
		"ltp_config_x86",
		"ltp_config_x86_64",
		"measure_io_as_jar",
		"nos_app_avb_service_genc++",
		"nos_app_avb_service_genc++_headers",
		"nos_app_avb_service_genc++_mock",
		"nos_app_identity_service_genc++",
		"nos_app_identity_service_genc++_headers",
		"nos_app_identity_service_genc++_mock",
		"nos_app_keymaster_service_genc++",
		"nos_app_keymaster_service_genc++_headers",
		"nos_app_keymaster_service_genc++_mock",
		"nos_app_weaver_service_genc++",
		"nos_app_weaver_service_genc++_headers",
		"nos_app_weaver_service_genc++_mock",
		"nos_generator_test_service_genc++",
		"nos_generator_test_service_genc++_headers",
		"nos_generator_test_service_genc++_mock",
		"openwrt_rootfs_combined_aarch64",
		"openwrt_rootfs_combined_x86_64",
		"openwrt_rootfs_customization_aarch64",
		"openwrt_rootfs_customization_x86_64",
		"pandora-python-gen-src",
		"pdl_cxx_canonical_be_src_gen",
		"pdl_cxx_canonical_be_test_gen",
		"pdl_cxx_canonical_le_src_gen",
		"pdl_cxx_canonical_le_test_gen",
		"pdl_python_generator_be_test_gen",
		"pdl_python_generator_le_test_gen",
		"pdl_rust_noalloc_le_test_backend_srcs",
		"pdl_rust_noalloc_le_test_gen_harness",
		"pixelatoms_defs.h",
		"pixelstatsatoms.cpp",
		"pixelstatsatoms.h",
		"pvmfw_fdt_template_rs",
		"r8retrace-dexdump-sample-app",
		"r8retrace-run-retrace",
		"rootcanal_bredr_bb_packets_cxx_gen",
		"rootcanal_hci_packets_cxx_gen",
		"rootcanal_link_layer_packets_cxx_gen",
		"sample-profile-text-protos",
		"seller-frontend-service-stub-lite",
		"services.core.protologsrc",
		"statsd-config-protos",
		"swiftshader_spvtools_update_build_version",
		"temp_layoutlib",
		"ue_unittest_erofs_imgs",
		"uwb_core_artifacts",
		"vm-tests-tf-lib",
		"vndk_abi_dump_zip",
		"vts_vndk_abi_dump_zip",
		"wm_shell_protolog_src",
		"wmtests.protologsrc",
		// go/keep-sorted end
	}

	SandboxingDenyPathList = []string{
		// go/keep-sorted start
		"art/test",
		"external/cronet",
		// go/keep-sorted end
	}
)
