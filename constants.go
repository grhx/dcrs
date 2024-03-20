package discord

type Event string
type Opcode int
type Intent int
type UserFlag int
type UserPremium int
type MessageType int
type MessageFlag int

const (
	// events
	Event_READY                                  Event = "READY"
	Event_RESUMED                                Event = "RESUMED"
	Event_RECONNECT                              Event = "RECONNECT"
	Event_INVALID_SESSION                        Event = "INVALID_SESSION"
	Event_APPLICATION_COMMAND_PERMISSIONS_UPDATE Event = "APPLICATION_COMMAND_PERMISSIONS_UPDATE"
	Event_AUTO_MODERATION_RULE_CREATE            Event = "AUTO_MODERATION_RULE_CREATE"
	Event_AUTO_MODERATION_RULE_UPDATE            Event = "AUTO_MODERATION_RULE_UPDATE"
	Event_AUTO_MODERATION_RULE_DELETE            Event = "AUTO_MODERATION_RULE_DELETE"
	Event_AUTO_MODERATION_ACTION_EXECUTION       Event = "AUTO_MODERATION_ACTION_EXECUTION"
	Event_CHANNEL_CREATE                         Event = "CHANNEL_CREATE"
	Event_CHANNEL_UPDATE                         Event = "CHANNEL_UPDATE"
	Event_CHANNEL_DELETE                         Event = "CHANNEL_DELETE"
	Event_CHANNEL_PINS_UPDATE                    Event = "CHANNEL_PINS_UPDATE"
	Event_THREAD_CREATE                          Event = "THREAD_CREATE"
	Event_THREAD_UPDATE                          Event = "THREAD_UPDATE"
	Event_THREAD_DELETE                          Event = "THREAD_DELETE"
	Event_THREAD_LIST_SYNC                       Event = "THREAD_LIST_SYNC"
	Event_THREAD_MEMBER_UPDATE                   Event = "THREAD_MEMBER_UPDATE"
	Event_THREAD_MEMBERS_UPDATE                  Event = "THREAD_MEMBERS_UPDATE"
	Event_ENTITLEMENT_CREATE                     Event = "ENTITLEMENT_CREATE"
	Event_ENTITLEMENT_UPDATE                     Event = "ENTITLEMENT_UPDATE"
	Event_ENTITLEMENT_DELETE                     Event = "ENTITLEMENT_DELETE"
	Event_GUILD_CREATE                           Event = "GUILD_CREATE"
	Event_GUILD_UPDATE                           Event = "GUILD_UPDATE"
	Event_GUILD_DELETE                           Event = "GUILD_DELETE"
	Event_GUILD_AUDIT_LOG_ENTRY_CREATE           Event = "GUILD_AUDIT_LOG_ENTRY_CREATE"
	Event_GUILD_BAN_ADD                          Event = "GUILD_BAN_ADD"
	Event_GUILD_BAN_REMOVE                       Event = "GUILD_BAN_REMOVE"
	Event_GUILD_EMOJIS_UPDATE                    Event = "GUILD_EMOJIS_UPDATE"
	Event_GUILD_STICKERS_UPDATE                  Event = "GUILD_STICKERS_UPDATE"
	Event_GUILD_INTEGRATIONS_UPDATE              Event = "GUILD_INTEGRATIONS_UPDATE"
	Event_GUILD_MEMBER_ADD                       Event = "GUILD_MEMBER_ADD"
	Event_GUILD_MEMBER_REMOVE                    Event = "GUILD_MEMBER_REMOVE"
	Event_GUILD_MEMBER_UPDATE                    Event = "GUILD_MEMBER_UPDATE"
	Event_GUILD_MEMBER_CHUNK                     Event = "GUILD_MEMBER_CHUNK"
	Event_GUILD_ROLE_CREATE                      Event = "GUILD_ROLE_CREATE"
	Event_GUILD_ROLE_UPDATE                      Event = "GUILD_ROLE_UPDATE"
	Event_GUILD_ROLE_DELETE                      Event = "GUILD_ROLE_DELETE"
	Event_GUILD_SCHEDULED_EVENT_CREATE           Event = "GUILD_SCHEDULED_EVENT_CREATE"
	Event_GUILD_SCHEDULED_EVENT_UPDATE           Event = "GUILD_SCHEDULED_EVENT_UPDATE"
	Event_GUILD_SCHEDULED_EVENT_DELETE           Event = "GUILD_SCHEDULED_EVENT_DELETE"
	Event_GUILD_SCHEDULED_EVENT_USER_ADD         Event = "GUILD_SCHEDULED_EVENT_USER_ADD"
	Event_GUILD_SCHEDULED_EVENT_USER_REMOVE      Event = "GUILD_SCHEDULED_EVENT_USER_REMOVE"
	Event_INTEGRATION_CREATE                     Event = "INTEGRATION_CREATE"
	Event_INTEGRATION_UPDATE                     Event = "INTEGRATION_UPDATE"
	Event_INTEGRATION_DELETE                     Event = "INTEGRATION_DELETE"
	Event_INVITE_CREATE                          Event = "INVITE_CREATE"
	Event_INVITE_DELETE                          Event = "INVITE_DELETE"
	Event_MESSAGE_CREATE                         Event = "MESSAGE_CREATE"
	Event_MESSAGE_UPDATE                         Event = "MESSAGE_UPDATE"
	Event_MESSAGE_DELETE                         Event = "MESSAGE_DELETE"
	Event_MESSAGE_DELETE_BULK                    Event = "MESSAGE_DELETE_BULK"
	Event_MESSAGE_REACTION_ADD                   Event = "MESSAGE_REACTION_ADD"
	Event_MESSAGE_REACTION_REMOVE                Event = "MESSAGE_REACTION_REMOVE"
	Event_MESSAGE_REACTION_REMOVE_ALL            Event = "MESSAGE_REACTION_REMOVE_ALL"
	Event_MESSAGE_REACTION_REMOVE_EMOJI          Event = "MESSAGE_REACTION_REMOVE_EMOJI"
	Event_PRESENCE_UPDATE                        Event = "PRESENCE_UPDATE"
	Event_STAGE_INSTANCE_CREATE                  Event = "STAGE_INSTANCE_CREATE"
	Event_STAGE_INSTANCE_UPDATE                  Event = "STAGE_INSTANCE_UPDATE"
	Event_STAGE_INSTANCE_DELETE                  Event = "STAGE_INSTANCE_DELETE"
	Event_TYPING_START                           Event = "TYPING_START"
	Event_USER_UPDATE                            Event = "USER_UPDATE"
	Event_VOICE_STATE_UPDATE                     Event = "VOICE_STATE_UPDATE"
	Event_VOICE_SERVER_UPDATE                    Event = "VOICE_SERVER_UPDATE"
	Event_WEBHOOKS_UPDATE                        Event = "WEBHOOKS_UPDATE"
	// opcodes
	Opcode_DISPATCH              Opcode = 0
	Opcode_HEARTBEAT             Opcode = 1
	Opcode_IDENTIFY              Opcode = 2
	Opcode_PRESENCE_UPDATE       Opcode = 3
	Opcode_VOICE_STATE_UPDATE    Opcode = 4
	Opcode_RESUME                Opcode = 6
	Opcode_RECONNECT             Opcode = 7
	Opcode_REQUEST_GUILD_MEMBERS Opcode = 8
	Opcode_INVALID_SESSION       Opcode = 9
	Opcode_HELLO                 Opcode = 10
	Opcode_HEARTBEAT_ACK         Opcode = 11
	// intent
	Intent_ALL                           Intent = 3276799
	Intent_GUILDS                        Intent = 1 << 0
	Intent_GUILD_MEMBERS                 Intent = 1 << 1
	Intent_GUILD_MODERATION              Intent = 1 << 2
	Intent_GUILD_EMOJIS_AND_STICKERS     Intent = 1 << 3
	Intent_INTEGRATIONS                  Intent = 1 << 4
	Intent_WEBHOOKS                      Intent = 1 << 5
	Intent_INVITES                       Intent = 1 << 6
	Intent_GUILD_VOICE_STATES            Intent = 1 << 7
	Intent_GUILD_PRESENCES               Intent = 1 << 8
	Intent_GUILD_MESSAGES                Intent = 1 << 9
	Intent_GUILD_MESSAGE_REACTIONS       Intent = 1 << 10
	Intent_GUILD_MESSAGE_TYPING          Intent = 1 << 11
	Intent_DIRECT_MESSAGES               Intent = 1 << 12
	Intent_DIRECT_MESSAGE_REACTIONS      Intent = 1 << 13
	Intent_DIRECT_MESSAGE_TYPING         Intent = 1 << 14
	Intent_MESSAGE_CONTENT               Intent = 1 << 15
	Intent_GUILD_SCHEDULED_EVENTS        Intent = 1 << 16
	Intent_AUTO_MODERATION_CONFIGURATION Intent = 1 << 20
	Intent_AUTO_MODERATION_EXECUTION     Intent = 1 << 21
	// user flag
	UserFlag_STAFF                   UserFlag = 1 << 0
	UserFlag_PARTNER                 UserFlag = 1 << 1
	UserFlag_HYPESQUAD               UserFlag = 1 << 2
	UserFlag_BUG_HUNTER_LEVEL_1      UserFlag = 1 << 3
	UserFlag_HYPSQUAD_ONLINE_HOUSE_1 UserFlag = 1 << 6
	UserFlag_HYPSQUAD_ONLINE_HOUSE_2 UserFlag = 1 << 7
	UserFlag_HYPSQUAD_ONLINE_HOUSE_3 UserFlag = 1 << 8
	UserFlag_PREMIUM_EARLY_SUPPORTER UserFlag = 1 << 9
	UserFlag_TEAM_PSEUDO_USER        UserFlag = 1 << 10
	UserFlag_BUG_HUNTER_LEVEL_2      UserFlag = 1 << 14
	UserFlag_VERIFIED_BOT            UserFlag = 1 << 16
	UserFlag_VERIFIED_DEVELOPER      UserFlag = 1 << 17
	UserFlag_CERTIFIED_MODERATOR     UserFlag = 1 << 18
	UserFlag_BOT_HTTP_INTERACTIONS   UserFlag = 1 << 19
	UserFlag_ACTIVE_DEVELOPER        UserFlag = 1 << 22
	// user premiu
	UserPremium_NONE          UserPremium = 0
	UserPremium_NITRO_CLASSIC UserPremium = 1
	UserPremium_NITRO         UserPremium = 2
	UserPremium_NITRO_BASIC   UserPremium = 3
	// message typ
	MessageType_DEFAULT                                      MessageType = 0
	MessageType_RECIPIENT_ADD                                MessageType = 1
	MessageType_RECIPIENT_REMOVE                             MessageType = 2
	MessageType_CALL                                         MessageType = 3
	MessageType_CHANNEL_NAME_CHANGE                          MessageType = 4
	MessageType_CHANNEL_ICON_CHANGE                          MessageType = 5
	MessageType_CHANNEL_PINNED_MESSAGE                       MessageType = 6
	MessageType_USER_JOIN                                    MessageType = 7
	MessageType_GUILD_BOOST                                  MessageType = 8
	MessageType_GUILD_BOOST_TIER_1                           MessageType = 9
	MessageType_GUILD_BOOST_TIER_2                           MessageType = 10
	MessageType_GUILD_BOOST_TIER_3                           MessageType = 11
	MessageType_CHANNEL_FOLLOW_ADD                           MessageType = 12
	MessageType_GUILD_DISCOVERY_DISQUALIFIED                 MessageType = 14 // skip 13 lol
	MessageType_GUILD_DISCOVERY_REQUALIFIED                  MessageType = 15
	MessageType_GUILD_DISCOVERY_GRACE_PERIOD_INITIAL_WARNING MessageType = 16
	MessageType_GUILD_DISCOVERY_GRACE_PERIOD_FINAL_WARNING   MessageType = 17
	MessageType_THREAD_CREATED                               MessageType = 18
	MessageType_REPLY                                        MessageType = 19
	MessageType_CHAT_INPUT_COMMAND                           MessageType = 20
	MessageType_THREAD_STARTER_MESSAGE                       MessageType = 21
	MessageType_GUILD_INVITE_REMINDER                        MessageType = 22
	MessageType_CONTEXT_MENU_COMMAND                         MessageType = 23
	MessageType_AUTO_MODERATION_ACTION                       MessageType = 24
	MessageType_ROLE_SUBSCRIPTION_PURCHASE                   MessageType = 25
	MessageType_INTERACTION_PREMIUM_UPSELL                   MessageType = 26
	MessageType_STAGE_START                                  MessageType = 27
	MessageType_STAGE_END                                    MessageType = 28
	MessageType_STAGE_SPEAKER                                MessageType = 29
	MessageType_STAGE_TOPIC                                  MessageType = 31
	MessageType_GUILD_APPLICATION_PREMIUM_SUBSCRIPTION       MessageType = 32
	// message flag
	MessageFlag_CROSSPOSTED                            MessageFlag = 1 << 0
	MessageFlag_IS_CROSSPOST                           MessageFlag = 1 << 1
	MessageFlag_SUPPRESS_EMBEDS                        MessageFlag = 1 << 2
	MessageFlag_SOURCE_MESSAGE_DELETED                 MessageFlag = 1 << 3
	MessageFlag_URGENT                                 MessageFlag = 1 << 4
	MessageFlag_HAS_THREAD                             MessageFlag = 1 << 5
	MessageFlag_EPHEMERAL                              MessageFlag = 1 << 6
	MessageFlag_LOADING                                MessageFlag = 1 << 7
	MessageFlag_FAILED_TO_MENTION_SOME_ROLES_IN_THREAD MessageFlag = 1 << 8
	MessageFlag_SUPPRESS_NOTIFICATIONS                 MessageFlag = 1 << 12
	MessageFlag_IS_VOICE_MESSAGE                       MessageFlag = 1 << 13
)
