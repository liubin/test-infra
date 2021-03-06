package githubql

// CommentAuthorAssociation represents a comment author association with repository.
type CommentAuthorAssociation string

// A comment author association with repository.
const (
	CommentAuthorAssociationMember               CommentAuthorAssociation = "MEMBER"                 // Author is a member of the organization that owns the repository.
	CommentAuthorAssociationOwner                CommentAuthorAssociation = "OWNER"                  // Author is the owner of the repository.
	CommentAuthorAssociationCollaborator         CommentAuthorAssociation = "COLLABORATOR"           // Author has been invited to collaborate on the repository.
	CommentAuthorAssociationContributor          CommentAuthorAssociation = "CONTRIBUTOR"            // Author has previously committed to the repository.
	CommentAuthorAssociationFirstTimeContributor CommentAuthorAssociation = "FIRST_TIME_CONTRIBUTOR" // Author has not previously committed to the repository.
	CommentAuthorAssociationFirstTimer           CommentAuthorAssociation = "FIRST_TIMER"            // Author has not previously committed to GitHub.
	CommentAuthorAssociationNone                 CommentAuthorAssociation = "NONE"                   // Author has no association with the repository.
)

// CommentCannotUpdateReason represents the possible errors that will prevent a user from updating a comment.
type CommentCannotUpdateReason string

// The possible errors that will prevent a user from updating a comment.
const (
	CommentCannotUpdateReasonInsufficientAccess    CommentCannotUpdateReason = "INSUFFICIENT_ACCESS"     // You must be the author or have write access to this repository to update this comment.
	CommentCannotUpdateReasonLocked                CommentCannotUpdateReason = "LOCKED"                  // Unable to create comment because issue is locked.
	CommentCannotUpdateReasonLoginRequired         CommentCannotUpdateReason = "LOGIN_REQUIRED"          // You must be logged in to update this comment.
	CommentCannotUpdateReasonMaintenance           CommentCannotUpdateReason = "MAINTENANCE"             // Repository is under maintenance.
	CommentCannotUpdateReasonVerifiedEmailRequired CommentCannotUpdateReason = "VERIFIED_EMAIL_REQUIRED" // At least one email address must be verified to update this comment.
)

// DefaultRepositoryPermissionField represents the possible default permissions for organization-owned repositories.
type DefaultRepositoryPermissionField string

// The possible default permissions for organization-owned repositories.
const (
	DefaultRepositoryPermissionFieldRead  DefaultRepositoryPermissionField = "READ"  // Members have read access to org repos by default.
	DefaultRepositoryPermissionFieldWrite DefaultRepositoryPermissionField = "WRITE" // Members have read and write access to org repos by default.
	DefaultRepositoryPermissionFieldAdmin DefaultRepositoryPermissionField = "ADMIN" // Members have read, write, and admin access to org repos by default.
)

// DeploymentState represents the possible states in which a deployment can be.
type DeploymentState string

// The possible states in which a deployment can be.
const (
	DeploymentStateAbandoned DeploymentState = "ABANDONED" // The pending deployment was not updated after 30 minutes.
	DeploymentStateActive    DeploymentState = "ACTIVE"    // The deployment is currently active.
	DeploymentStateDestroyed DeploymentState = "DESTROYED" // An inactive transient deployment.
	DeploymentStateError     DeploymentState = "ERROR"     // The deployment experienced an error.
	DeploymentStateFailure   DeploymentState = "FAILURE"   // The deployment has failed.
	DeploymentStateInactive  DeploymentState = "INACTIVE"  // The deployment is inactive.
	DeploymentStatePending   DeploymentState = "PENDING"   // The deployment is pending.
)

// DeploymentStatusState represents the possible states for a deployment status.
type DeploymentStatusState string

// The possible states for a deployment status.
const (
	DeploymentStatusStatePending  DeploymentStatusState = "PENDING"  // The deployment is pending.
	DeploymentStatusStateSuccess  DeploymentStatusState = "SUCCESS"  // The deployment was successful.
	DeploymentStatusStateFailure  DeploymentStatusState = "FAILURE"  // The deployment has failed.
	DeploymentStatusStateInactive DeploymentStatusState = "INACTIVE" // The deployment is inactive.
	DeploymentStatusStateError    DeploymentStatusState = "ERROR"    // The deployment experienced an error.
)

// GistOrderField represents properties by which gist connections can be ordered.
type GistOrderField string

// Properties by which gist connections can be ordered.
const (
	GistOrderFieldCreatedAt GistOrderField = "CREATED_AT" // Order gists by creation time.
	GistOrderFieldUpdatedAt GistOrderField = "UPDATED_AT" // Order gists by update time.
	GistOrderFieldPushedAt  GistOrderField = "PUSHED_AT"  // Order gists by push time.
)

// GistPrivacy represents the privacy of a Gist.
type GistPrivacy string

// The privacy of a Gist.
const (
	GistPrivacyPublic GistPrivacy = "PUBLIC" // Public.
	GistPrivacySecret GistPrivacy = "SECRET" // Secret.
	GistPrivacyAll    GistPrivacy = "ALL"    // Gists that are public and secret.
)

// GitSignatureState represents the state of a Git signature.
type GitSignatureState string

// The state of a Git signature.
const (
	GitSignatureStateValid                GitSignatureState = "VALID"                 // Valid signature and verified by GitHub.
	GitSignatureStateInvalid              GitSignatureState = "INVALID"               // Invalid signature.
	GitSignatureStateMalformedSig         GitSignatureState = "MALFORMED_SIG"         // Malformed signature.
	GitSignatureStateUnknownKey           GitSignatureState = "UNKNOWN_KEY"           // Key used for signing not known to GitHub.
	GitSignatureStateBadEmail             GitSignatureState = "BAD_EMAIL"             // Invalid email used for signing.
	GitSignatureStateUnverifiedEmail      GitSignatureState = "UNVERIFIED_EMAIL"      // Email used for signing unverified on GitHub.
	GitSignatureStateNoUser               GitSignatureState = "NO_USER"               // Email used for signing not known to GitHub.
	GitSignatureStateUnknownSigType       GitSignatureState = "UNKNOWN_SIG_TYPE"      // Unknown signature type.
	GitSignatureStateUnsigned             GitSignatureState = "UNSIGNED"              // Unsigned.
	GitSignatureStateGpgverifyUnavailable GitSignatureState = "GPGVERIFY_UNAVAILABLE" // Internal error - the GPG verification service is unavailable at the moment.
	GitSignatureStateGpgverifyError       GitSignatureState = "GPGVERIFY_ERROR"       // Internal error - the GPG verification service misbehaved.
	GitSignatureStateNotSigningKey        GitSignatureState = "NOT_SIGNING_KEY"       // The usage flags for the key that signed this don't allow signing.
	GitSignatureStateExpiredKey           GitSignatureState = "EXPIRED_KEY"           // Signing key expired.
)

// IssueOrderField represents properties by which issue connections can be ordered.
type IssueOrderField string

// Properties by which issue connections can be ordered.
const (
	IssueOrderFieldCreatedAt IssueOrderField = "CREATED_AT" // Order issues by creation time.
	IssueOrderFieldUpdatedAt IssueOrderField = "UPDATED_AT" // Order issues by update time.
	IssueOrderFieldComments  IssueOrderField = "COMMENTS"   // Order issues by comment count.
)

// IssuePubSubTopic represents the possible PubSub channels for an issue.
type IssuePubSubTopic string

// The possible PubSub channels for an issue.
const (
	IssuePubSubTopicUpdated    IssuePubSubTopic = "UPDATED"    // The channel ID for observing issue updates.
	IssuePubSubTopicMarkasread IssuePubSubTopic = "MARKASREAD" // The channel ID for marking an issue as read.
)

// IssueState represents the possible states of an issue.
type IssueState string

// The possible states of an issue.
const (
	IssueStateOpen   IssueState = "OPEN"   // An issue that is still open.
	IssueStateClosed IssueState = "CLOSED" // An issue that has been closed.
)

// LanguageOrderField represents properties by which language connections can be ordered.
type LanguageOrderField string

// Properties by which language connections can be ordered.
const (
	LanguageOrderFieldSize LanguageOrderField = "SIZE" // Order languages by the size of all files containing the language.
)

// MergeableState represents whether or not a PullRequest can be merged.
type MergeableState string

// Whether or not a PullRequest can be merged.
const (
	MergeableStateMergeable   MergeableState = "MERGEABLE"   // The pull request can be merged.
	MergeableStateConflicting MergeableState = "CONFLICTING" // The pull request cannot be merged due to merge conflicts.
	MergeableStateUnknown     MergeableState = "UNKNOWN"     // The mergeability of the pull request is still being calculated.
)

// MilestoneState represents the possible states of a milestone.
type MilestoneState string

// The possible states of a milestone.
const (
	MilestoneStateOpen   MilestoneState = "OPEN"   // A milestone that is still open.
	MilestoneStateClosed MilestoneState = "CLOSED" // A milestone that has been closed.
)

// OrderDirection represents possible directions in which to order a list of items when provided an `orderBy` argument.
type OrderDirection string

// Possible directions in which to order a list of items when provided an `orderBy` argument.
const (
	OrderDirectionAsc  OrderDirection = "ASC"  // Specifies an ascending order for a given `orderBy` argument.
	OrderDirectionDesc OrderDirection = "DESC" // Specifies a descending order for a given `orderBy` argument.
)

// OrganizationInvitationRole represents the possible organization invitation roles.
type OrganizationInvitationRole string

// The possible organization invitation roles.
const (
	OrganizationInvitationRoleDirectMember   OrganizationInvitationRole = "DIRECT_MEMBER"   // The user is invited to be a direct member of the organization.
	OrganizationInvitationRoleAdmin          OrganizationInvitationRole = "ADMIN"           // The user is invited to be an admin of the organization.
	OrganizationInvitationRoleBillingManager OrganizationInvitationRole = "BILLING_MANAGER" // The user is invited to be a billing manager of the organization.
	OrganizationInvitationRoleReinstate      OrganizationInvitationRole = "REINSTATE"       // The user's previous role will be reinstated.
)

// ProjectCardState represents various content states of a ProjectCard.
type ProjectCardState string

// Various content states of a ProjectCard.
const (
	ProjectCardStateContentOnly ProjectCardState = "CONTENT_ONLY" // The card has content only.
	ProjectCardStateNoteOnly    ProjectCardState = "NOTE_ONLY"    // The card has a note only.
	ProjectCardStateRedacted    ProjectCardState = "REDACTED"     // The card is redacted.
)

// ProjectOrderField represents properties by which project connections can be ordered.
type ProjectOrderField string

// Properties by which project connections can be ordered.
const (
	ProjectOrderFieldCreatedAt ProjectOrderField = "CREATED_AT" // Order projects by creation time.
	ProjectOrderFieldUpdatedAt ProjectOrderField = "UPDATED_AT" // Order projects by update time.
	ProjectOrderFieldName      ProjectOrderField = "NAME"       // Order projects by name.
)

// ProjectState represents state of the project; either 'open' or 'closed'.
type ProjectState string

// State of the project; either 'open' or 'closed'.
const (
	ProjectStateOpen   ProjectState = "OPEN"   // The project is open.
	ProjectStateClosed ProjectState = "CLOSED" // The project is closed.
)

// PullRequestPubSubTopic represents the possible PubSub channels for a pull request.
type PullRequestPubSubTopic string

// The possible PubSub channels for a pull request.
const (
	PullRequestPubSubTopicUpdated    PullRequestPubSubTopic = "UPDATED"    // The channel ID for observing pull request updates.
	PullRequestPubSubTopicMarkasread PullRequestPubSubTopic = "MARKASREAD" // The channel ID for marking an pull request as read.
	PullRequestPubSubTopicHeadRef    PullRequestPubSubTopic = "HEAD_REF"   // The channel ID for observing head ref updates.
)

// PullRequestReviewEvent represents the possible events to perform on a pull request review.
type PullRequestReviewEvent string

// The possible events to perform on a pull request review.
const (
	PullRequestReviewEventComment        PullRequestReviewEvent = "COMMENT"         // Submit general feedback without explicit approval.
	PullRequestReviewEventApprove        PullRequestReviewEvent = "APPROVE"         // Submit feedback and approve merging these changes.
	PullRequestReviewEventRequestChanges PullRequestReviewEvent = "REQUEST_CHANGES" // Submit feedback that must be addressed before merging.
	PullRequestReviewEventDismiss        PullRequestReviewEvent = "DISMISS"         // Dismiss review so it now longer effects merging.
)

// PullRequestReviewState represents the possible states of a pull request review.
type PullRequestReviewState string

// The possible states of a pull request review.
const (
	PullRequestReviewStatePending          PullRequestReviewState = "PENDING"           // A review that has not yet been submitted.
	PullRequestReviewStateCommented        PullRequestReviewState = "COMMENTED"         // An informational review.
	PullRequestReviewStateApproved         PullRequestReviewState = "APPROVED"          // A review allowing the pull request to merge.
	PullRequestReviewStateChangesRequested PullRequestReviewState = "CHANGES_REQUESTED" // A review blocking the pull request from merging.
	PullRequestReviewStateDismissed        PullRequestReviewState = "DISMISSED"         // A review that has been dismissed.
)

// PullRequestState represents the possible states of a pull request.
type PullRequestState string

// The possible states of a pull request.
const (
	PullRequestStateOpen   PullRequestState = "OPEN"   // A pull request that is still open.
	PullRequestStateClosed PullRequestState = "CLOSED" // A pull request that has been closed without being merged.
	PullRequestStateMerged PullRequestState = "MERGED" // A pull request that has been closed by being merged.
)

// ReactionContent represents emojis that can be attached to Issues, Pull Requests and Comments.
type ReactionContent string

// Emojis that can be attached to Issues, Pull Requests and Comments.
const (
	ReactionContentThumbsUp   ReactionContent = "THUMBS_UP"   // Represents the 👍 emoji.
	ReactionContentThumbsDown ReactionContent = "THUMBS_DOWN" // Represents the 👎 emoji.
	ReactionContentLaugh      ReactionContent = "LAUGH"       // Represents the 😄 emoji.
	ReactionContentHooray     ReactionContent = "HOORAY"      // Represents the 🎉 emoji.
	ReactionContentConfused   ReactionContent = "CONFUSED"    // Represents the 😕 emoji.
	ReactionContentHeart      ReactionContent = "HEART"       // Represents the ❤️ emoji.
)

// ReactionOrderField represents a list of fields that reactions can be ordered by.
type ReactionOrderField string

// A list of fields that reactions can be ordered by.
const (
	ReactionOrderFieldCreatedAt ReactionOrderField = "CREATED_AT" // Allows ordering a list of reactions by when they were created.
)

// RepositoryAffiliation represents the affiliation of a user to a repository.
type RepositoryAffiliation string

// The affiliation of a user to a repository.
const (
	RepositoryAffiliationOwner              RepositoryAffiliation = "OWNER"               // Repositories that are owned by the authenticated user.
	RepositoryAffiliationCollaborator       RepositoryAffiliation = "COLLABORATOR"        // Repositories that the user has been added to as a collaborator.
	RepositoryAffiliationOrganizationMember RepositoryAffiliation = "ORGANIZATION_MEMBER" // Repositories that the user has access to through being a member of an organization. This includes every repository on every team that the user is on.
)

// RepositoryCollaboratorAffiliation represents the affiliation type between collaborator and repository.
type RepositoryCollaboratorAffiliation string

// The affiliation type between collaborator and repository.
const (
	RepositoryCollaboratorAffiliationAll     RepositoryCollaboratorAffiliation = "ALL"     // All collaborators of the repository.
	RepositoryCollaboratorAffiliationOutside RepositoryCollaboratorAffiliation = "OUTSIDE" // All outside collaborators of an organization-owned repository.
)

// RepositoryLockReason represents the possible reasons a given repository could be in a locked state.
type RepositoryLockReason string

// The possible reasons a given repository could be in a locked state.
const (
	RepositoryLockReasonMoving    RepositoryLockReason = "MOVING"    // The repository is locked due to a move.
	RepositoryLockReasonBilling   RepositoryLockReason = "BILLING"   // The repository is locked due to a billing related reason.
	RepositoryLockReasonRename    RepositoryLockReason = "RENAME"    // The repository is locked due to a rename.
	RepositoryLockReasonMigrating RepositoryLockReason = "MIGRATING" // The repository is locked due to a migration.
)

// RepositoryOrderField represents properties by which repository connections can be ordered.
type RepositoryOrderField string

// Properties by which repository connections can be ordered.
const (
	RepositoryOrderFieldCreatedAt  RepositoryOrderField = "CREATED_AT" // Order repositories by creation time.
	RepositoryOrderFieldUpdatedAt  RepositoryOrderField = "UPDATED_AT" // Order repositories by update time.
	RepositoryOrderFieldPushedAt   RepositoryOrderField = "PUSHED_AT"  // Order repositories by push time.
	RepositoryOrderFieldName       RepositoryOrderField = "NAME"       // Order repositories by name.
	RepositoryOrderFieldStargazers RepositoryOrderField = "STARGAZERS" // Order repositories by number of stargazers.
)

// RepositoryPermission represents the access level to a repository.
type RepositoryPermission string

// The access level to a repository.
const (
	RepositoryPermissionAdmin RepositoryPermission = "ADMIN" // Can read, clone, push, and add collaborators.
	RepositoryPermissionWrite RepositoryPermission = "WRITE" // Can read, clone and push.
	RepositoryPermissionRead  RepositoryPermission = "READ"  // Can read and clone.
)

// RepositoryPrivacy represents the privacy of a repository.
type RepositoryPrivacy string

// The privacy of a repository.
const (
	RepositoryPrivacyPublic  RepositoryPrivacy = "PUBLIC"  // Public.
	RepositoryPrivacyPrivate RepositoryPrivacy = "PRIVATE" // Private.
)

// SearchType represents represents the individual results of a search.
type SearchType string

// Represents the individual results of a search.
const (
	SearchTypeIssue      SearchType = "ISSUE"      // Returns results matching issues in repositories.
	SearchTypeRepository SearchType = "REPOSITORY" // Returns results matching repositories.
	SearchTypeUser       SearchType = "USER"       // Returns results matching users on GitHub.
)

// StarOrderField represents properties by which star connections can be ordered.
type StarOrderField string

// Properties by which star connections can be ordered.
const (
	StarOrderFieldStarredAt StarOrderField = "STARRED_AT" // Allows ordering a list of stars by when they were created.
)

// StatusState represents the possible commit status states.
type StatusState string

// The possible commit status states.
const (
	StatusStateExpected StatusState = "EXPECTED" // Status is expected.
	StatusStateError    StatusState = "ERROR"    // Status is errored.
	StatusStateFailure  StatusState = "FAILURE"  // Status is failing.
	StatusStatePending  StatusState = "PENDING"  // Status is pending.
	StatusStateSuccess  StatusState = "SUCCESS"  // Status is successful.
)

// SubscriptionState represents the possible states of a subscription.
type SubscriptionState string

// The possible states of a subscription.
const (
	SubscriptionStateUnsubscribed SubscriptionState = "UNSUBSCRIBED" // The User is only notified when particpating or @mentioned.
	SubscriptionStateSubscribed   SubscriptionState = "SUBSCRIBED"   // The User is notified of all conversations.
	SubscriptionStateIgnored      SubscriptionState = "IGNORED"      // The User is never notified.
)

// TeamMemberRole represents the possible team member roles; either 'maintainer' or 'member'.
type TeamMemberRole string

// The possible team member roles; either 'maintainer' or 'member'.
const (
	TeamMemberRoleMaintainer TeamMemberRole = "MAINTAINER" // A team maintainer has permission to add and remove team members.
	TeamMemberRoleMember     TeamMemberRole = "MEMBER"     // A team member has no administrative permissions on the team.
)

// TeamMembershipType represents defines which types of team members are included in the returned list. Can be one of IMMEDIATE, CHILD_TEAM or ALL.
type TeamMembershipType string

// Defines which types of team members are included in the returned list. Can be one of IMMEDIATE, CHILD_TEAM or ALL.
const (
	TeamMembershipTypeImmediate TeamMembershipType = "IMMEDIATE"  // Includes only immediate members of the team.
	TeamMembershipTypeChildTeam TeamMembershipType = "CHILD_TEAM" // Includes only child team members for the team.
	TeamMembershipTypeAll       TeamMembershipType = "ALL"        // Includes immediate and child team members for the team.
)

// TeamOrderField represents properties by which team connections can be ordered.
type TeamOrderField string

// Properties by which team connections can be ordered.
const (
	TeamOrderFieldName TeamOrderField = "NAME" // Allows ordering a list of teams by name.
)

// TeamPrivacy represents the possible team privacy values.
type TeamPrivacy string

// The possible team privacy values.
const (
	TeamPrivacySecret  TeamPrivacy = "SECRET"  // A secret team can only be seen by its members.
	TeamPrivacyVisible TeamPrivacy = "VISIBLE" // A visible team can be seen and @mentioned by every member of the organization.
)

// TeamRepositoryOrderField represents properties by which team repository connections can be ordered.
type TeamRepositoryOrderField string

// Properties by which team repository connections can be ordered.
const (
	TeamRepositoryOrderFieldCreatedAt  TeamRepositoryOrderField = "CREATED_AT" // Order repositories by creation time.
	TeamRepositoryOrderFieldUpdatedAt  TeamRepositoryOrderField = "UPDATED_AT" // Order repositories by update time.
	TeamRepositoryOrderFieldPushedAt   TeamRepositoryOrderField = "PUSHED_AT"  // Order repositories by push time.
	TeamRepositoryOrderFieldName       TeamRepositoryOrderField = "NAME"       // Order repositories by name.
	TeamRepositoryOrderFieldPermission TeamRepositoryOrderField = "PERMISSION" // Order repositories by permission.
	TeamRepositoryOrderFieldStargazers TeamRepositoryOrderField = "STARGAZERS" // Order repositories by number of stargazers.
)

// TeamRole represents the role of a user on a team.
type TeamRole string

// The role of a user on a team.
const (
	TeamRoleAdmin  TeamRole = "ADMIN"  // User has admin rights on the team.
	TeamRoleMember TeamRole = "MEMBER" // User is a member of the team.
)

// TopicSuggestionDeclineReason represents reason that the suggested topic is declined.
type TopicSuggestionDeclineReason string

// Reason that the suggested topic is declined.
const (
	TopicSuggestionDeclineReasonNotRelevant        TopicSuggestionDeclineReason = "NOT_RELEVANT"        // The suggested topic is not relevant to the repository.
	TopicSuggestionDeclineReasonTooSpecific        TopicSuggestionDeclineReason = "TOO_SPECIFIC"        // The suggested topic is too specific for the repository (e.g. #ruby-on-rails-version-4-2-1).
	TopicSuggestionDeclineReasonPersonalPreference TopicSuggestionDeclineReason = "PERSONAL_PREFERENCE" // The viewer does not like the suggested topic.
	TopicSuggestionDeclineReasonTooGeneral         TopicSuggestionDeclineReason = "TOO_GENERAL"         // The suggested topic is too general for the repository.
)
