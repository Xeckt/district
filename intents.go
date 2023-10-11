package main

import "github.com/bwmarrin/discordgo"

func IntentString(i discordgo.Intent) string {
	switch i {
	case discordgo.IntentsGuilds:
		return "Guilds"
	case discordgo.IntentGuildMembers:
		return "GuildMembers"
	case discordgo.IntentsGuildBans:
		return "GuildBans"
	case discordgo.IntentGuildEmojis:
		return "GuildEmojis"
	case discordgo.IntentGuildIntegrations:
		return "GuildIntegrations"
	case discordgo.IntentGuildWebhooks:
		return "GuildWebhooks"
	case discordgo.IntentGuildInvites:
		return "GuildInvites"
	case discordgo.IntentGuildVoiceStates:
		return "GuildVoiceStates"
	case discordgo.IntentGuildPresences:
		return "GuildPresences"
	case discordgo.IntentGuildMessages:
		return "GuildMessages"
	case discordgo.IntentGuildMessageReactions:
		return "GuildMessageREactions"
	case discordgo.IntentGuildMessageTyping:
		return "GuildMessageTyping"
	case discordgo.IntentDirectMessages:
		return "DirectMessages"
	case discordgo.IntentDirectMessageReactions:
		return "DirectMessageReactions"
	case discordgo.IntentDirectMessageTyping:
		return "DirectMessageTyping"
	case discordgo.IntentMessageContent:
		return "MessageContent"
	case discordgo.IntentGuildScheduledEvents:
		return "GuildScheduledEvents"
	case discordgo.IntentAutoModerationConfiguration:
		return "AutoModerationConfiguration"
	case discordgo.IntentAutoModerationExecution:
		return "AutoModerationExecution"
	default:
		return "No string found for this intent"
	}
}

func SpecifyIntents(im HandlerManager) {
	im.AddIntents(
		discordgo.IntentGuilds,
		discordgo.IntentGuildMembers,
		discordgo.IntentGuildMessages,
	)
}
