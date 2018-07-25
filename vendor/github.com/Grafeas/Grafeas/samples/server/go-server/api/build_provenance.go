/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * OpenAPI spec version: 0.1
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

// Provenance of a build. Contains all information needed to verify the full details about the build from source to completion.
type BuildProvenance struct {
	// Unique identifier of the build.
	Id string `json:"id,omitempty"`

	// ID of the project.
	ProjectId string `json:"projectId,omitempty"`

	// Numerical ID of the project.
	ProjectNum string `json:"projectNum,omitempty"`

	// Commands requested by the build.
	Commands []Command `json:"commands,omitempty"`

	// Output of the build.
	BuiltArtifacts []Artifact `json:"builtArtifacts,omitempty"`

	// Time at which the build was created.
	CreateTime string `json:"createTime,omitempty"`

	// Time at which execution of the build was started.
	StartTime string `json:"startTime,omitempty"`

	// Time at whihc execution of the build was finished.
	FinishTime string `json:"finishTime,omitempty"`

	// GAIA ID of end user who initiated this build; at the time that the BuildProvenance is uploaded to Analysis, this will be resolved to the primary e-mail address of the user and stored in the Creator field.
	UserId string `json:"userId,omitempty"`

	// E-mail address of the user who initiated this build. Note that this was the user's e-mail address at the time the build was initiated; this address may not represent the same end-user for all time.
	Creator string `json:"creator,omitempty"`

	// Google Cloud Storage bucket where logs were written.
	LogsBucket string `json:"logsBucket,omitempty"`

	// Details of the Source input to the build.
	SourceProvenance Source `json:"sourceProvenance,omitempty"`

	// Trigger identifier if the build was triggered automatically; empty if not.
	TriggerId string `json:"triggerId,omitempty"`

	// Special options applied to this build. This is a catch-all field where build providers can enter any desired additional details.
	BuildOptions map[string]string `json:"buildOptions,omitempty"`

	// Version string of the builder at the time this build was executed.
	BuilderVersion string `json:"builderVersion,omitempty"`
}
