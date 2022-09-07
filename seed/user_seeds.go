package seed

import (
	"context"
	"education/config"
	"fmt"
)

func Do(ctx context.Context) error {

	_, err := config.Client.User.
		Create().
		SetName("Mahmoud").
		SetChannelName("Droos").
		SetEmail("mhdshaikh20403@gmail.com").
		SetSpecialist("Feqh").
		SetAge(26).
		SetPhone("01152569522").
		SetLanguage("Arabic").
		SetCountry("Cairo, Egypt").
		SetShorBio("Channel To Teach You the Feqh Subject.").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User: %w", err)
	}

	_, err = config.Client.User.
		Create().
		SetName("Mohammed").
		SetChannelName("Sona Channel").
		SetEmail("sonah@gmail.com").
		SetSpecialist("Sona").
		SetAge(39).
		SetPhone("01254869812").
		SetLanguage("Arabic").
		SetCountry("Alex, Egypt").
		SetShorBio("Channel To Teach You the Sona Subject.").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User: %w", err)
	}

	_, err = config.Client.User.
		Create().
		SetName("Ibrahim").
		SetChannelName("Akeda Channel").
		SetEmail("akeda@gmail.com").
		SetSpecialist("Akeda").
		SetAge(22).
		SetPhone("01587945632").
		SetLanguage("Arabic").
		SetCountry("Alsharkeya, Egypt").
		SetShorBio("Channel To Teach You the Akeda Subject.").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User: %w", err)
	}

	_, err = config.Client.User.
		Create().
		SetName("Sayed").
		SetChannelName("Nahw Channel").
		SetEmail("nahw@gmail.com").
		SetSpecialist("Feqh").
		SetAge(29).
		SetPhone("01547856985").
		SetLanguage("Arabic").
		SetCountry("Elmonofia, Egypt").
		SetShorBio("Channel To Teach You the Nahw Subject.").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User: %w", err)
	}

	_, err = config.Client.User.
		Create().
		SetName("Ahmed").
		SetChannelName("Qouran").
		SetEmail("qouran@gmail.com").
		SetSpecialist("Qouran").
		SetAge(40).
		SetPhone("01078469532").
		SetLanguage("Arabic").
		SetCountry("Cairo, Egypt").
		SetShorBio("Channel To Teach You the Qouran.").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User: %w", err)
	}

	_, err = config.Client.User.
		Create().
		SetName("Abdo").
		SetChannelName("General").
		SetEmail("general@gmail.com").
		SetSpecialist("General").
		SetAge(36).
		SetPhone("01247189635").
		SetLanguage("Arabic").
		SetCountry("Cairo, Egypt").
		SetShorBio("Channel To Teach You the Genaral Subjects.").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User: %w", err)
	}

	_, err = config.Client.User.
		Create().
		SetName("Ali").
		SetChannelName("Elsharh").
		SetEmail("aliabdo@gmail.com").
		SetSpecialist("General").
		SetAge(33).
		SetPhone("01184596214").
		SetLanguage("Arabic").
		SetCountry("Cairo, Egypt").
		SetShorBio("Channel To Teach You the Feqh Subject.").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User: %w", err)
	}

	_, err = config.Client.User.
		Create().
		SetName("Mahmoud").
		SetChannelName("Droos Online").
		SetEmail("droosonline@gmail.com").
		SetSpecialist("General").
		SetAge(50).
		SetPhone("0125874696").
		SetLanguage("Arabic").
		SetCountry("Elgharbia, Egypt").
		SetShorBio("Channel To Teach You the General Subjects.").
		Save(ctx)

	if err != nil {
		return fmt.Errorf("creating User: %w", err)
	}
	return nil
}
