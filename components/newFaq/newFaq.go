package newFaq

import "log"

const DATA_KEY = "faqData"
const CONVERTER_FUNCTION_NAME = "cast_to_faq"

type NewFaq struct {
	Question string
	Answer   string
}

func Cast_to_faq(mapper map[string]interface{}) []NewFaq {
	var v []NewFaq
	var ok bool
	if x, found := mapper[DATA_KEY]; found {
		v, ok = x.([]NewFaq)
		if !ok {
			log.Fatal("Unable to case....")
		}
		return v
	}
	log.Fatal("Not found in the map")
	return v
}

func NewDummy() []NewFaq {
	return []NewFaq{
		{
			Question: "What is this platform, and how can it help my business?",
			Answer:   "It’s a SaaS solution to streamline workflows, boost collaboration, and drive results.",
		},
		{
			Question: "Is there a free trial available?",
			Answer:   "Yes, enjoy a 14-day free trial with full feature access.",
		},
		{
			Question: "Do I need technical expertise to use this software?",
			Answer:   "No, it’s user-friendly and includes tutorials for easy onboarding.",
		},
		{
			Question: "Can I integrate this with other tools I already use?",
			Answer:   "Yes, it integrates with tools like Slack, Google Workspace, and Zapier.",
		},
		{
			Question: "What kind of support can I expect?",
			Answer:   "We offer 24/7 support via chat, email, and phone, plus a knowledge base.",
		},
		{
			Question: "How secure is my data on your platform?",
			Answer:   "Your data is protected with encryption and compliance with GDPR and SOC 2.",
		},
	}

}
