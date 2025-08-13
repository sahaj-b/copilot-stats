package main

import (
	"fmt"
	"strings"
	"time"
)

func calculateUsage(entitlement, remaining int) (int, float64) {
	used := entitlement - remaining
	var usagePercent float64
	if entitlement > 0 {
		usagePercent = float64(used) / float64(entitlement) * 100
	}
	return used, usagePercent
}

func makeProgressBar(percent float64, width int) string {
	filled := int(float64(width) * percent / 100)
	filledBar := strings.Repeat(barChar, filled)
	emptyBar := strings.Repeat(barChar, width-filled)

	// Color the progress bar based on percentage
	var color string
	switch {
	case percent >= 80:
		color = red
	case percent >= 60:
		color = yellow
	case percent >= 40:
		color = yellow
	default:
		color = green
	}

	if gray == "" { // NO_COLOR active (colors disabled)
		emptyBar = strings.Repeat(" ", width-filled)
	}

	return color + filledBar + gray + emptyBar + reset
}

func displayPremiumInteractions(premium *struct {
	Entitlement      int     `json:"entitlement"`
	Remaining        int     `json:"remaining"`
	PercentRemaining float64 `json:"percent_remaining"`
	Unlimited        bool    `json:"unlimited"`
	OveragePermitted bool    `json:"overage_permitted"`
},
) {
	fmt.Println(bold + yellow + "🚀 Premium Interactions" + reset)

	used, usagePercent := calculateUsage(premium.Entitlement, premium.Remaining)

	bar := makeProgressBar(usagePercent, 20)
	// Progress bar with percent in plain/bold only
	fmt.Println("   " + bar + " " + bold + fmt.Sprintf("%.1f%%", usagePercent) + reset)

	// Labels bold blue, values plain (bold only where needed)
	fmt.Println("   " + bold + blue + "Used:" + reset + "      " + bold + fmt.Sprint(used) + reset + " / " + bold + fmt.Sprint(premium.Entitlement) + reset)

	fmt.Println("   " + bold + blue + "Remaining:" + reset + " " + fmt.Sprint(premium.Remaining) + reset)

	statusText := "Unlimited"
	statusColor := green
	if !premium.Unlimited {
		statusText = "Limited"
		statusColor = yellow
	}
	fmt.Println("   " + bold + blue + "Status:" + reset + "    " + statusColor + statusText + reset)

	overageText := "Permitted"
	overageColor := green
	if !premium.OveragePermitted {
		overageText = "Not Permitted"
		overageColor = red
	}
	fmt.Println("   " + bold + blue + "Overage:" + reset + "   " + overageColor + overageText + reset)

	fmt.Println()
}

func displayChat(chat *struct {
	Entitlement int  `json:"entitlement"`
	Remaining   int  `json:"remaining"`
	Unlimited   bool `json:"unlimited"`
},
) {
	fmt.Println(bold + yellow + "💬 Chat" + reset)

	if chat.Unlimited {
		fmt.Println("   " + bold + blue + "Status:" + reset + " " + green + "Unlimited" + reset)
	} else {
		used, usagePercent := calculateUsage(chat.Entitlement, chat.Remaining)
		// Progress bar with percent
		bar := makeProgressBar(usagePercent, 20)
		fmt.Println("   " + bar + " " + bold + fmt.Sprintf("%.1f%%", usagePercent) + reset)

		fmt.Println("   " + bold + blue + "Used:" + reset + " " + bold + fmt.Sprint(used) + reset + " / " + bold + fmt.Sprint(chat.Entitlement) + reset)
		fmt.Println("   " + bold + blue + "Remaining:" + reset + " " + fmt.Sprint(chat.Remaining) + reset)
	}
	fmt.Println()
}

func displayCompletions(completions *struct {
	Entitlement int  `json:"entitlement"`
	Remaining   int  `json:"remaining"`
	Unlimited   bool `json:"unlimited"`
},
) {
	fmt.Println(bold + yellow + "🔧 Completions" + reset)

	if completions.Unlimited {
		fmt.Println("   " + bold + blue + "Status:" + reset + " " + green + "Unlimited" + reset)
	} else {
		used, usagePercent := calculateUsage(completions.Entitlement, completions.Remaining)
		// Progress bar with percent
		bar := makeProgressBar(usagePercent, 20)
		fmt.Println("   " + bar + " " + bold + fmt.Sprintf("%.1f%%", usagePercent) + reset)

		fmt.Println("   " + bold + blue + "Used:" + reset + " " + bold + fmt.Sprint(used) + reset + " / " + bold + fmt.Sprint(completions.Entitlement) + reset)
		fmt.Println("   " + bold + blue + "Remaining:" + reset + " " + fmt.Sprint(completions.Remaining) + reset)
	}
	fmt.Println()
}

func displayQuotaReset(resetDate string) {
	if resetDate == "" {
		return
	}

	fmt.Println(bold + yellow + "⏰ Quota Reset Information" + reset)

	if t, err := time.Parse("2006-01-02", resetDate); err == nil {
		resetDate = t.Format("Jan 2, 2006")
	}

	resetTime, err := time.Parse("2006-01-02", resetDate)
	if err != nil {
		resetTime, _ = time.Parse("Jan 2, 2006", resetDate)
	}
	if !resetTime.IsZero() {
		daysLeft := int(time.Until(resetTime).Hours() / 24)
		percent := float64(30-daysLeft) / 30 * 100
		if percent < 0 {
			percent = 0
		} else if percent > 100 {
			percent = 100
		}

		bar := makeProgressBar(percent, 20)
		fmt.Println("   " + bar + reset)
		fmt.Println("   " + bold + blue + "Days left:" + reset + " " + fmt.Sprint(daysLeft))
		fmt.Println("   " + bold + blue + "Resets on:" + reset + " " + fmt.Sprint(resetDate))

	}
	fmt.Println()
}

func displayCopilotStats(stats *CopilotStats) {
	if stats.QuotaSnapshots.PremiumInteractions != nil {
		displayPremiumInteractions(stats.QuotaSnapshots.PremiumInteractions)
	}

	displayQuotaReset(stats.QuotaResetDate)

	if stats.QuotaSnapshots.Chat != nil {
		displayChat(stats.QuotaSnapshots.Chat)
	}

	if stats.QuotaSnapshots.Completions != nil {
		displayCompletions(stats.QuotaSnapshots.Completions)
	}
}
