package validation

//func TestIfFirstTrueContinueToSecond(t *testing.T) {
//	validPipline := WrapFunc("valid", func(signedMessage *message.SignedMessage) error {
//		return nil
//	})
//
//	invalidPipeline := WrapFunc("invalid", func(signedMessage *message.SignedMessage) error {
//		return errors.Errorf("error")
//	})
//
//	t.Run("a valid, b valid", func(t *testing.T) {
//		require.NoError(t, IfFirstTrueContinueToSecond(validPipline, validPipline).Run(nil))
//	})
//
//	t.Run("a valid, b invalid", func(t *testing.T) {
//		require.EqualError(t, IfFirstTrueContinueToSecond(validPipline, invalidPipeline).Run(nil), "error")
//	})
//
//	t.Run("a invalid, b valid", func(t *testing.T) {
//		require.NoError(t, IfFirstTrueContinueToSecond(invalidPipeline, validPipline).Run(nil))
//	})
//
//	t.Run("a invalid, b invalid", func(t *testing.T) {
//		require.NoError(t, IfFirstTrueContinueToSecond(invalidPipeline, invalidPipeline).Run(nil))
//	})
//}
